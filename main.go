package main

import (
	"context"
	"fmt"
	"os"

	"github.com/normalbeans/ladder/ladder"
	"github.com/normalbeans/ladder/ladder/component"
	"github.com/normalbeans/ladder/ladder/key"
	"github.com/normalbeans/ladder/ladder/state"
	"golang.org/x/term"
)

func main() {

	originalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), originalState)

	ctx, cancel := context.WithCancel(context.Background())
	l := ladder.Ladder{
		Components: make(map[int]component.Component),
		State: state.LState{
			CompModels: make(map[int]component.Model),
			Focus:      0,
			WIDTH:      100,
			HEIGHT:     20,
			CURSORX:    1,
			CURSORY:    1,
			Changed:    make(map[int]bool),
		},
		Quit:       make(chan byte, 1),
		OldState:   originalState,
		Input:      make(chan key.Key),
		Background: make(chan component.BackgroundMsg),
		Ctx:        ctx,
		Cancel:     cancel,
	}

	b := &component.Box{}
	bModel := b.Init(component.BoxModel{
		Width:   30,
		Height:  5,
		Originx: 1,
		Originy: 1,
	})
	l.RegisterComponent(b, bModel)

	c := &component.Counter{}
	cModel := c.Init(component.CounterModel{
		Width:   40,
		Height:  3,
		Originx: 1,
		Originy: 6,
	})
	l.RegisterComponent(c, cModel)

	LEGEND := &component.Legend{}
	LEGENDMODEL := LEGEND.Init(component.LegendModel{
		Width:   100,
		Height:  1,
		Originx: 1,
		Originy: 12,
	})
	l.RegisterComponent(LEGEND, LEGENDMODEL)

	go l.Snooper()
	go l.Looper()

	<-l.Quit
	term.Restore(int(os.Stdin.Fd()), l.OldState)
	fmt.Print("\x1b[H\x1b[J\x1b[H\x1b[?25h")
}
