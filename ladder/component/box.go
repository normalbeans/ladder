package component

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"strings"
)

type BoxModel struct {
	Width, Height    int
	Originx, Originy int
	Color            int
	Controls         Command
}

type Box struct {
	id int
}

func getRandomColor() int {
	return 31 + rand.IntN(5)
}

// func (b *Box) IncreaseWidth(state LState) LState {
// 	currentModel := state.CompModels[b.id].(BoxModel)
// 	currentModel.Width += 1
// 	state.CompModels[b.id] = currentModel
// 	return state
// }

func (b BoxModel) GetControls() Command {
	return b.Controls
}

// Implement interface

func (b Box) GetID() int {
	return b.id
}

func (b *Box) SetID(id int) {
	b.id = id
}

func (b *Box) Render(state LState) {
	c := state.CompModels[b.id].(BoxModel)
	fmt.Printf("\x1b[%d;%dH", c.Originy, c.Originx)
	for j := 0; j < c.Height; j++ {
		switch j {
		case 0:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2580", c.Width-2)+"\u2588\x1b[0m", c.Color, c.Originy+j, c.Originx)
		case c.Height - 1:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2584", c.Width-2)+"\u2588\x1b[0m", c.Color, c.Originy+j, c.Originx)
		default:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat(" ", c.Width-2)+"\u2588\x1b[0m", c.Color, c.Originy+j, c.Originx)
		}
	}
}

// not part of interface but kind of required.

func (b *Box) DataModel(width, height, ox, oy int, controls Command) Model {

	defaultControls := map[string]ComponentFunction{
		"c": {
			KeyHint: "c",
			Legend:  "Randomise color",
			Function: func(state LState) LState {
				cstate := state.CompModels[b.GetID()].(BoxModel)
				cstate.Color = 31 + rand.IntN(5)
				state.CompModels[b.GetID()] = cstate
				return state
			},
		},
	}

	maps.Copy(defaultControls, controls.Actions)

	return BoxModel{
		Width:   width,
		Height:  height,
		Originx: ox,
		Originy: oy,
		Color:   getRandomColor(),
		Controls: Command{
			Actions: defaultControls,
		},
	}
}
