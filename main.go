package main

import (
	"os"

	"github.com/nn-advith/ladder/ladder"
	"golang.org/x/term"
)

func main() {

	nl := ladder.InitLadder() // determine the structure here

	// ladder.InitLadder(ladder.FLow{
	// 	ScreenOptions: {
	// 		//
	// 	},
	// 	PageOptions: {
	// 		//
	// 	},
	// 	COmponentst: ///
	// })

	defer term.Restore(int(os.Stdin.Fd()), nl.TermState)

	go nl.Render()
	ladder.InputListener()

}
