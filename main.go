package main

import (
	"fmt"
	"os"

	"github.com/nn-advith/ladder/ladder"
	"golang.org/x/term"
)

func main() {

	originalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), originalState)

	b := ladder.Box{
		Width:   30,
		Height:  5,
		Originx: 1,
		Originy: 1,
	}
	b2 := ladder.Box{
		Width:   60,
		Height:  3,
		Originx: 1,
		Originy: 6,
	}

	b.SetColor()

	l := ladder.Ladder{
		Components: []ladder.Component{&b, &b2},
		Focus:      0,
		CURSORX:    1,
		CURSORY:    1,
		WIDTH:      100,
		HEIGHT:     20,
		Data: map[int]ladder.Model{
			0: b.DataModel(),
		},
		Quit:     make(chan byte, 1),
		OldState: originalState,
	}

	go l.Snooper()
	l.Render()
	<-l.Quit
	term.Restore(int(os.Stdin.Fd()), l.OldState)
	fmt.Print("\x1b[H\x1b[J\x1b[H\x1b[?25h")
}
