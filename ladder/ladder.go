package ladder

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/normalbeans/ladder/ladder/component"
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
	l.Components[newID] = c
}

func (l *Ladder) Render() {
	fmt.Print("\x1b[H\x1b[J\x1b[H\x1b[?25l")
	for i := 0; i < len(l.Components); i++ {
		l.Components[i].Render(l.State)
	}
	var focusindex strings.Builder
	for k, v := range l.Components[l.State.Focus].Controls().Actions {
		fmt.Fprintf(&focusindex, " %s - %s |", k, v.Legend)
	}
	fmt.Printf("\x1b[%d;%dH"+focusindex.String(), l.State.HEIGHT+1, 1)
}

func (l *Ladder) Looper() {
	for {
		select {
		case <-l.Quit:
			return
		default:
			l.Render()
			time.Sleep(16 * time.Millisecond)
		}
	}
}

func (l *Ladder) Snooper() {
	for {
		buf := make([]byte, 8)
		n, err := os.Stdin.Read(buf)
		data := buf[:n]
		if err != nil {
			panic(err)
		}
		if data[0] == 0x03 {
			close(l.Quit)
		} else if data[0] == 27 && data[1] == 91 {
			switch data[2] {
			case 65:
				l.State.Focus = (l.State.Focus - 1 + len(l.State.CompModels)) % len(l.State.CompModels)
			case 66:
				l.State.Focus = (l.State.Focus + 1) % len(l.State.CompModels)
			}
		} else {
			if executor, ok := l.Components[l.State.Focus].Controls().Actions[string(data)]; ok {
				l.State = executor.Function(l.State)
			}
		}
	}
}
