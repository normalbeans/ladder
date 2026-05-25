package ladder

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/normalbeans/ladder/ladder/component"
	"github.com/normalbeans/ladder/ladder/key"
	"github.com/normalbeans/ladder/ladder/state"
	"golang.org/x/term"
)

type Msg interface{}

// Ladder
type Ladder struct {
	State      state.LState
	Components map[int]component.Component
	OldState   *term.State
	Quit       chan byte
	Input      chan key.Key
	Background chan component.BackgroundMsg
	Ctx        context.Context
	Cancel     context.CancelFunc
}

func (l *Ladder) RegisterComponent(c component.Component, m component.Model) {
	newID := len(l.State.CompModels)
	c.SetID(newID)
	l.State.CompModels[newID] = m
	l.State.Changed[newID] = true
	l.Components[newID] = c

	for _, f := range m.GetBackgroundFunc() {
		fn := f
		go fn(l.Ctx, newID, l.Background)
	}
}

func (l *Ladder) Render() {
	for i := 0; i < len(l.Components); i++ {
		if l.State.Changed[i] || l.State.CompModels[i].GetReRenderPolicy() == component.ReRenderAlways {
			l.Components[i].Render(l.State.CompModels[l.State.Focus])
			l.State.Changed[i] = false
		}
	}
}

func (l *Ladder) Looper() {
	if len(l.Components) == 0 {
		l.Cancel()
		close(l.Quit)
	}
	fmt.Print("\x1b[H\x1b[J\x1b[?25l")
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-l.Quit:
			return
		case data := <-l.Input:
			switch data {
			case key.CtrlC, key.Escape:
				l.Cancel()
				close(l.Quit)
			case key.ArrowUp:
				l.State.Focus = (l.State.Focus - 1 + len(l.State.CompModels)) % len(l.State.CompModels)
				// rerender = true
			case key.ArrowDown:
				l.State.Focus = (l.State.Focus + 1) % len(l.State.CompModels)
				// rerender = true
			default:
				if executor, ok := l.State.CompModels[l.State.Focus].GetControls()[string(data)]; ok {
					actionresult := executor.Function(nil)
					nmodel, changed := l.State.CompModels[l.State.Focus].Update(l.State.CompModels[l.State.Focus], actionresult)
					if changed {
						l.State.CompModels[l.State.Focus] = nmodel
						l.State.Changed[l.State.Focus] = true
					}
					// rerender = true
				}
			}
		case data := <-l.Background:
			nm, updated := l.State.CompModels[data.ID].Update(l.State.CompModels[data.ID], data.Data)
			if updated {
				l.State.CompModels[data.ID] = nm
				l.State.Changed[data.ID] = true
			}
		case <-ticker.C:
			l.Render()
		}
	}

}

func (l *Ladder) Snooper() {

	// l.Render()
	for {
		// rerender := false
		buf := make([]byte, 8)
		n, err := os.Stdin.Read(buf)
		data := string(buf[:n])
		if err != nil {
			panic(err)
		}
		l.Input <- key.Key(data)
		// if rerender {
		// 	l.Render()
		// }
	}
}
