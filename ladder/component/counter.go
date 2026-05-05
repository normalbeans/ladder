package component

import "fmt"

type CounterModel struct {
	Width, Height, Originx, Originy int
	Count                           int
}

type Counter struct {
	id int
}

func (c *Counter) Increment(state LState) LState {
	currentModel := state.CompModels[c.id].(CounterModel)
	currentModel.Count++
	state.CompModels[c.id] = currentModel
	return state
}

func (c *Counter) Decrement(state LState) LState {
	currentModel := state.CompModels[c.id].(CounterModel)
	currentModel.Count--
	state.CompModels[c.id] = currentModel
	return state
}

// Implement interface

func (c *Counter) SetID(newID int) {
	c.id = newID
}

func (c *Counter) Render(state LState) {
	m := state.CompModels[c.id].(CounterModel)
	fmt.Printf("\x1b[%d;%dH", m.Originy, m.Originx)
	for i := 0; i < m.Height; i++ {
		if i == m.Height/2 {
			fmt.Printf("\x1b[%d;%dH\tCount: %d", m.Originy+i, 1, m.Count)
		} else {
			fmt.Printf("\x1b[%d;%dH", m.Originy+i, 1)
		}
	}
}

func (c *Counter) Controls() Command {
	return Command{
		Actions: map[string]ComponentFunction{
			"+": {
				KeyHint:  "+",
				Legend:   "Increment",
				Function: c.Increment,
			},
			"-": {
				KeyHint:  "-",
				Legend:   "Decrement",
				Function: c.Decrement,
			},
		},
	}
}

func (c *Counter) DataModel(width, height, ox, oy int) Model {
	return CounterModel{
		Width:   width,
		Height:  height,
		Originx: ox,
		Originy: oy,
		Count:   0,
	}
}
