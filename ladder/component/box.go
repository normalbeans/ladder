package component

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/normalbeans/ladder/ladder/key"
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

func (b *Box) IncreaseWidth(state LState) LState {
	currentModel := state.CompModels[b.id].(BoxModel)
	currentModel.Width += 1
	state.CompModels[b.id] = currentModel
	return state
}

// CORE

func (b *Box) SetID(id int) {
	b.id = id
}

func (b *Box) Render(state LState) {
	c := state.CompModels[b.id].(BoxModel)
	fmt.Printf("\x1b[%d;%dH", c.Originy, c.Originx)
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
				KeyHint:  "c",
				Legend:   "Change Color",
				Function: b.SetColor,
			},
			key.ArrowRight: {
				KeyHint:  "->",
				Legend:   "Increase width",
				Function: b.IncreaseWidth,
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
