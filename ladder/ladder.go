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
	// DependencyChanged chan int
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (l *Ladder) RegisterComponent(c component.Component, m component.Model) int {
	newID := len(l.State.CompModels)
	c.SetID(newID)
	l.State.CompModels[newID] = m
	l.State.Changed[newID] = true

	mDependencies := m.GetDependencies()
	for _, v := range mDependencies {
		l.State.DependencyMap[v.Id] = append(l.State.DependencyMap[v.Id], newID)
	}
	l.Components[newID] = c

	for _, f := range m.GetBackgroundFunc() {
		fn := f
		go fn(l.Ctx, newID, l.Background)
	}

	return newID
}

func (l *Ladder) Render() {
	for i := 0; i < len(l.Components); i++ {
		if l.State.Changed[i] || l.State.CompModels[i].GetReRenderPolicy() == component.ReRenderAlways {

			l.Components[i].Render(component.RenderContext{
				SelfModel:  l.State.CompModels[i],
				FocusModel: l.State.CompModels[l.State.Focus],
			})
			l.State.Changed[i] = false
		}
	}
}

func (l *Ladder) updateDependencies(id int) {
	// for dependencies of id, update
	for _, v := range l.State.DependencyMap[id] {
		updateContext := component.UpdateContext{
			SelfModel: l.State.CompModels[v],
			Read: func(id int) component.Model {
				return l.State.CompModels[id]
			},
			Data: l.State.CompModels[v].GetData(),
		}
		nm, updated := l.State.CompModels[v].Update(updateContext)
		if updated {
			l.State.CompModels[v] = nm
			l.State.Changed[v] = true
			l.updateDependencies(v)
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
				n := len(l.State.CompModels)
				for range n {
					l.State.Focus = (l.State.Focus - 1 + n) % n
					if l.State.CompModels[l.State.Focus].IsFocusable() != component.FocusFalse {
						break
					}
				}
				// rerender = true
			case key.ArrowDown:
				// go to next focusable component
				n := len(l.State.CompModels)
				for range n {
					l.State.Focus = (l.State.Focus + 1) % n
					if l.State.CompModels[l.State.Focus].IsFocusable() != component.FocusFalse {
						break
					}
				}
				// rerender = true
			default:
				if executor, ok := l.State.CompModels[l.State.Focus].GetControls()[string(data)]; ok {

					updateContext := component.UpdateContext{
						SelfModel: l.State.CompModels[l.State.Focus],
						Read: func(id int) component.Model {
							return l.State.CompModels[id]
						},
						KeyBinding: string(data),
						Data:       nil,
					}

					updateData, err := executor.Function(updateContext)
					if err == nil {
						updateContext.Data = updateData
						nmodel, changed := l.State.CompModels[l.State.Focus].Update(updateContext)
						if changed {
							l.State.CompModels[l.State.Focus] = nmodel
							l.State.Changed[l.State.Focus] = true // here send info to reupdate dependencies
							l.updateDependencies(l.State.Focus)
						}
					}
					// rerender = true
				}
			}
		case data := <-l.Background:
			updateContext := component.UpdateContext{
				SelfModel: l.State.CompModels[data.ID],
				Read: func(id int) component.Model {
					return l.State.CompModels[id]
				},
				Data: data.Data,
			}
			nm, updated := l.State.CompModels[data.ID].Update(updateContext)
			if updated {
				l.State.CompModels[data.ID] = nm
				l.State.Changed[data.ID] = true
				l.updateDependencies(data.ID)
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
