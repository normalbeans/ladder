package ladder

import (
	"fmt"
	"os"
	"time"

	"github.com/normalbeans/ladder/ladder/component"
	"github.com/normalbeans/ladder/ladder/key"
	"golang.org/x/term"
)

type Msg interface{}

// Ladder
type Ladder struct {
	State      component.LState
	Components map[int]component.Component
	OldState   *term.State
	Quit       chan byte
	Input      chan key.Key
}

func (l *Ladder) RegisterComponent(c component.Component, m component.Model) {
	newID := len(l.State.CompModels)
	c.SetID(newID)
	l.State.CompModels[newID] = m
	l.State.Changed[newID] = true
	l.Components[newID] = c
}

func (l *Ladder) Render() {

	for i := 0; i < len(l.Components); i++ {
		if l.State.Changed[i] || l.State.CompModels[i].GetRenderControls() == component.RenderAlways {
			l.Components[i].Render(l.State)
			l.State.Changed[i] = false
		}
	}
	// probably move this into its own component
	// var focusindex strings.Builder
	// for _, v := range l.State.CompModels[l.State.Focus].GetControls() {
	// 	fmt.Fprintf(&focusindex, " %s - %s |", v.KeyHint, v.Legend)
	// }
	// fmt.Printf("\x1b[%d;%dH\x1b[0K"+focusindex.String(), l.State.HEIGHT+1, 1)
}

func (l *Ladder) Looper() {
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
				close(l.Quit)
			case key.ArrowUp:
				l.State.Focus = (l.State.Focus - 1 + len(l.State.CompModels)) % len(l.State.CompModels)
				// rerender = true
			case key.ArrowDown:
				l.State.Focus = (l.State.Focus + 1) % len(l.State.CompModels)
				// rerender = true
			default:
				if executor, ok := l.State.CompModels[l.State.Focus].GetControls()[string(data)]; ok {
					l.State = executor.Function(l.State)
					l.State.Changed[l.State.Focus] = true
					// rerender = true
				}
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
