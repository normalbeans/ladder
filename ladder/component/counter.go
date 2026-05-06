package component

import (
	"fmt"
	"maps"
)

type CounterModel struct {
	Width, Height, Originx, Originy int
	Count                           int
	Controls                        Command
}

func (c CounterModel) GetControls() Command {
	return c.Controls
}

type Counter struct {
	id int
}

// Implement interface
func (c Counter) GetID() int {
	return c.id
}

func (c *Counter) SetID(newID int) {
	c.id = newID
}

func (c *Counter) Render(state LState) {
	m := state.CompModels[c.id].(CounterModel)
	fmt.Printf("\x1b[%d;%dH", m.Originy, m.Originx)
	for i := 0; i < m.Height; i++ {
		if i == m.Height/2 {
			fmt.Printf("\x1b[%d;%dH\x1b[0K\tCount: %d", m.Originy+i, 1, m.Count)
		} else {
			fmt.Printf("\x1b[%d;%dH", m.Originy+i, 1)
		}
	}
}


// not part of interface -> change this 

func (c *Counter) DataModel(width, height, ox, oy int, controls Command) Model {

	defaultControls := map[string]ComponentFunction{
		"+": {
			KeyHint: "+",
			Legend:  "Increment",
			Function: func(state LState) LState {
				cstate := state.CompModels[c.GetID()].(CounterModel)
				cstate.Count++
				state.CompModels[c.GetID()] = cstate
				return state
			},
		},
		"-": {
			KeyHint: "-",
			Legend:  "Decrement",
			Function: func(state LState) LState {
				cstate := state.CompModels[c.GetID()].(CounterModel)
				cstate.Count--
				state.CompModels[c.GetID()] = cstate
				return state
			},
		},
		"z": {
			KeyHint: "z",
			Legend:  "Reset to 0",
			Function: func(state LState) LState {
				cstate := state.CompModels[c.GetID()].(CounterModel)
				cstate.Count = 0
				state.CompModels[c.GetID()] = cstate
				return state
			},
		},
	}

	maps.Copy(defaultControls, controls.Actions)
	return CounterModel{
		Width:   width,
		Height:  height,
		Originx: ox,
		Originy: oy,
		Count:   0,
		Controls: Command{
			Actions: defaultControls,
		},
	}
}
