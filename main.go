package main

import (
	"os"

	"github.com/nn-advith/ladder/ladder"
	"github.com/nn-advith/ladder/ladder/state"
	"golang.org/x/term"
)

func main() {

	renderchannel := ladder.CreateRenderChannel()

	nl := ladder.InitLadder()

	defer func() {
		if r := recover(); r != nil {
			term.Restore(int(os.Stdin.Fd()), nl.TermState)
		} else {
			term.Restore(int(os.Stdin.Fd()), nl.TermState)
		}

	}()

	go nl.Render(renderchannel)
	buf := make([]byte, 3)
	for {
		for i := range buf {
			buf[i] = 0
		}
		_, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		switch buf[0] {
		case 13, 10:
			// enter/return
			cs := state.GetState()
			renderchannel <- cs
		}
	}
}
