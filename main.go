package main

import (
	"fmt"
	"os"

	"github.com/normalbeans/ladder/ladder"
	"github.com/normalbeans/ladder/ladder/component"
	"golang.org/x/term"
)

func main() {

	originalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), originalState)

	l := ladder.Ladder{
		Components: make(map[int]component.Component),
		State: component.LState{
			CompModels: make(map[int]component.Model),
			Focus:      0,
			WIDTH:      100,
			HEIGHT:     20,
			CURSORX:    1,
			CURSORY:    1,
		},
		Quit:     make(chan byte, 1),
		OldState: originalState,
	}

	b := &component.Box{}
	bModel := b.DataModel(30, 5, 1, 1)
	l.RegisterComponent(b, bModel)
	b2 := &component.Box{}
	b2Model := b2.DataModel(60, 3, 1, 6)
	l.RegisterComponent(b2, b2Model)

	go l.Snooper()
	go l.Looper()
	<-l.Quit
	term.Restore(int(os.Stdin.Fd()), l.OldState)
	fmt.Print("\x1b[H\x1b[J\x1b[H\x1b[?25h")
}
