package component

import (
	"fmt"
	"maps"
)

var defaultCounterModel = CounterModel{
	Width:   40,
	Height:  1,
	Originx: 1,
	Originy: 1,
	Data: Data{
		"count": 0,
	},
	Controls: Command{
		"+": {
			KeyHint: "+",
			Legend:  "Increment",
		},
		"-": {
			KeyHint: "-",
			Legend:  "Decrement",
		},
		"z": {
			KeyHint: "z",
			Legend:  "Reset to 0",
		},
	},
	ReRenderPolicy: ReRenderOnChange,
	Focusable:      FocusTrue,
	Dependents:     []int{},
}

type CounterModel struct {
	Width, Height, Originx, Originy int
	Controls                        Command
	Data                            Data
	ReRenderPolicy                  ReRenderPolicy
	Focusable                       Focusable
	Dependents                      []int
}

func (c CounterModel) GetControls() Command {
	return c.Controls
}

func (c CounterModel) GetReRenderPolicy() ReRenderPolicy {
	return c.ReRenderPolicy
}

func (c CounterModel) IsFocusable() Focusable {
	return c.Focusable
}

func (c CounterModel) GetDependents() []int {
	return c.Dependents
}

func (c CounterModel) GetBackgroundFunc() []BackgroundFunc {
	return nil
}

func (c CounterModel) Update(u UpdateContext) (Model, bool) {
	self, ok := u.SelfModel.(CounterModel)
	if !ok {
		return c, false
	}
	if u.KeyBinding != "" {
		switch u.KeyBinding {
		case "+":
			self.Data["count"] = self.Data["count"].(int) + 1
		case "-":
			self.Data["count"] = self.Data["count"].(int) - 1
		case "z":
			self.Data["count"] = 0
		default:
			return self, false
		}
	}
	if u.Data != nil {
		c, ok := u.Data.(int)
		if !ok {
			return self, false
		}
		self.Data["count"] = c
	}
	return self, true
}

func (c CounterModel) usingDefault(m Model) Model {
	userModel := m.(CounterModel)

	if userModel.Width == 0 {
		userModel.Width = defaultCounterModel.Width
	}
	if userModel.Height == 0 {
		userModel.Height = defaultCounterModel.Height
	}
	if userModel.Originx == 0 {
		userModel.Originx = defaultCounterModel.Originx
	}
	if userModel.Originy == 0 {
		userModel.Originy = defaultCounterModel.Originy
	}

	if userModel.Data == nil {
		userModel.Data = defaultCounterModel.Data
	} else {
		tempData := make(Data)
		maps.Copy(tempData, defaultCounterModel.Data)
		maps.Copy(tempData, userModel.Data)
		userModel.Data = tempData
	}

	if userModel.Controls == nil {
		userModel.Controls = defaultCounterModel.Controls
	} else {
		tempControls := make(Command)
		maps.Copy(tempControls, defaultCounterModel.Controls)
		maps.Copy(tempControls, userModel.Controls)
		userModel.Controls = tempControls
	}

	if userModel.ReRenderPolicy == ReRenderUnset {
		userModel.ReRenderPolicy = defaultCounterModel.ReRenderPolicy
	}

	if userModel.Focusable == FocusUnset {
		userModel.Focusable = defaultCounterModel.Focusable
	}

	if userModel.Dependents == nil {
		userModel.Dependents = defaultCounterModel.Dependents
	}

	return userModel
}

// renderer
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

func (c *Counter) Render(r RenderContext) {
	m := r.SelfModel.(CounterModel)
	fmt.Printf("\x1b[%d;%dH", m.Originy, m.Originx)
	for i := 0; i < m.Height; i++ {
		if i == m.Height/2 {
			fmt.Printf("\x1b[%d;%dH\x1b[0K\tCount: %d", m.Originy+i, 1, m.Data["count"])
		} else {
			fmt.Printf("\x1b[%d;%dH", m.Originy+i, 1)
		}
	}
}

func (c *Counter) Init(m Model) Model {

	newBoxModel := defaultCounterModel.usingDefault(m)

	return newBoxModel
}

// // not part of interface -> change this

// func (c *Counter) DataModel(width, height, ox, oy int, data Data, controls Command) Model {

// 	defaultData := Data{
// 		"count": 0,
// 	}

// 	defaultControls := Command{
// "+": {
// 	KeyHint: "+",
// 	Legend:  "Increment",
// 	Function: func(state LState) LState {
// 		cstate := state.CompModels[c.GetID()].(CounterModel)
// 		cstate.Data["count"]++
// 		state.CompModels[c.GetID()] = cstate
// 		return state
// 	},
// },
// "-": {
// 	KeyHint: "-",
// 	Legend:  "Decrement",
// 	Function: func(state LState) LState {
// 		cstate := state.CompModels[c.GetID()].(CounterModel)
// 		cstate.Data["count"]--
// 		state.CompModels[c.GetID()] = cstate
// 		return state
// 	},
// },
// "z": {
// 	KeyHint: "z",
// 	Legend:  "Reset to 0",
// 	Function: func(state LState) LState {
// 		cstate := state.CompModels[c.GetID()].(CounterModel)
// 		cstate.Data["count"] = 0
// 		state.CompModels[c.GetID()] = cstate
// 		return state
// 	},
// },
// 	}

// 	maps.Copy(defaultControls, controls)
// 	maps.Copy(defaultData, data)
// 	return CounterModel{
// 		Width:    width,
// 		Height:   height,
// 		Originx:  ox,
// 		Originy:  oy,
// 		Data:     defaultData,
// 		Controls: defaultControls,
// 	}
// }
