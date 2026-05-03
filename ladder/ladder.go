package ladder

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"golang.org/x/term"
)

type ComponentFunction struct {
	Legend   string
	Function func()
}

type Command struct {
	Actions map[string]ComponentFunction
}

type Model interface{}

type Component interface {
	GetOrigin() (int, int)
	Render()
	Controls() Command
	DataModel() Model
}

// Box
type Box struct {
	Width, Height    int
	Originx, Originy int
	color            int
}

func (b *Box) GetOrigin() (int, int) {
	return b.Originx, b.Originy
}

func (b *Box) SetColor() {
	b.color = 31 + rand.IntN(5)
	b.Render()
}

func (b *Box) Render() {
	fmt.Printf("\x1b[%d;%dH", b.Originx, b.Originy)
	for j := 0; j < b.Height; j++ {
		switch j {
		case 0:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2580", b.Width-2)+"\u2588\x1b[0m", b.color, b.Originy+j, b.Originx)
		case b.Height - 1:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2584", b.Width-2)+"\u2588\x1b[0m", b.color, b.Originy+j, b.Originx)
		default:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat(" ", b.Width-2)+"\u2588\x1b[0m", b.color, b.Originy+j, b.Originx)
		}
	}
}

func (b *Box) Controls() Command {
	return Command{
		Actions: map[string]ComponentFunction{
			"c": {
				Legend:   "Change Color",
				Function: b.SetColor,
			},
		},
	}
}

func (b *Box) DataModel() Model {
	return struct{}{}
}

// Ladder
type Ladder struct {
	WIDTH, HEIGHT, CURSORX, CURSORY int
	Components                      []Component
	Focus                           int
	Data                            map[int]Model
	OldState                        *term.State
	Quit                            chan byte
}

func (l *Ladder) Render() {
	fmt.Print("\x1b[H\x1b[J\x1b[H\x1b[?25l")
	for i := 0; i < len(l.Components); i++ {
		l.Components[i].Render()
	}
	var focusindex strings.Builder
	for k, v := range l.Components[l.Focus].Controls().Actions {
		fmt.Fprintf(&focusindex, " %s - %s |", k, v.Legend)
	}
	fmt.Printf("\x1b[%d;%dH"+focusindex.String(), l.HEIGHT+1, 1)

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
				l.Focus = (l.Focus - 1) % len(l.Components)
			case 66:
				l.Focus = (l.Focus + 1) % len(l.Components)
			}
		} else {
			if _, ok := l.Components[l.Focus].Controls().Actions[string(data)]; ok {
				l.Components[l.Focus].Controls().Actions[string(data)].Function()
			}
		}
	}
}
