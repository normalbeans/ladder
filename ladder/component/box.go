package component

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type BoxModel struct {
	Width, Height    int
	Originx, Originy int
	color            int
}

type Box struct {
	id int
}

func getRandomColor() int {
	return 31 + rand.IntN(5)
}

func (b *Box) SetColor(state LState) LState {
	currentModel := state.CompModels[b.id].(BoxModel)
	currentModel.color = getRandomColor()
	state.CompModels[b.id] = currentModel
	return state

}

// CORE

func (b *Box) SetID(id int) {
	b.id = id
}

func (b *Box) Render(state LState) {
	c := state.CompModels[b.id].(BoxModel)
	fmt.Printf("\x1b[%d;%dH", c.Originx, c.Originy)
	for j := 0; j < c.Height; j++ {
		switch j {
		case 0:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2580", c.Width-2)+"\u2588\x1b[0m", c.color, c.Originy+j, c.Originx)
		case c.Height - 1:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2584", c.Width-2)+"\u2588\x1b[0m", c.color, c.Originy+j, c.Originx)
		default:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat(" ", c.Width-2)+"\u2588\x1b[0m", c.color, c.Originy+j, c.Originx)
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

func (b *Box) DataModel(width, height, ox, oy int) Model {
	return BoxModel{
		Width:   width,
		Height:  height,
		Originx: ox,
		Originy: oy,
		color:   getRandomColor(),
	}
}
