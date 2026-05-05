package ladder

import (
	"fmt"
	"os"
	"strings"

	"github.com/normalbeans/ladder/ladder/component"
	"github.com/normalbeans/ladder/ladder/key"
	"golang.org/x/term"
)

// Ladder
type Ladder struct {
	State      component.LState
	Components map[int]component.Component
	OldState   *term.State
	Quit       chan byte
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
		if l.State.Changed[i] {
			l.Components[i].Render(l.State)
			l.State.Changed[i] = false
		}
	}
	// probably move this into its own component
	var focusindex strings.Builder
	for _, v := range l.Components[l.State.Focus].Controls().Actions {
		fmt.Fprintf(&focusindex, " %s - %s |", v.KeyHint, v.Legend)
	}
	fmt.Printf("\x1b[%d;%dH"+focusindex.String(), l.State.HEIGHT+1, 1)
}

// func (l *Ladder) Looper() {
// 	for {
// 		select {
// 		case <-l.Quit:
// 			return
// 		default:
// 			l.Render()
// 			time.Sleep(16 * time.Millisecond)
// 		}
// 	}
// }

func (l *Ladder) Snooper() {
	fmt.Print("\x1b[H\x1b[J\x1b[?25l")
	for {
		l.Render()
		buf := make([]byte, 8)
		n, err := os.Stdin.Read(buf)
		data := string(buf[:n])
		if err != nil {
			panic(err)
		}
		switch data {
		case key.CtrlC, key.Escape:
			close(l.Quit)
		case key.ArrowUp:
			l.State.Focus = (l.State.Focus - 1 + len(l.State.CompModels)) % len(l.State.CompModels)
		case key.ArrowDown:
			l.State.Focus = (l.State.Focus + 1) % len(l.State.CompModels)
		default:
			if executor, ok := l.Components[l.State.Focus].Controls().Actions[string(data)]; ok {
				l.State = executor.Function(l.State)
				l.State.Changed[l.State.Focus] = true
			}
		}

	}
}
