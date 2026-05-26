package component

import (
	"errors"
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
			Function: func(uc UpdateContext) (any, error) {
				self, ok := uc.SelfModel.(CounterModel)
				if !ok {
					return nil, errors.New("ERROR: Assertion failed")
				}
				return self.Data["count"].(int) + 1, nil
			},
		},
		"-": {
			KeyHint: "-",
			Legend:  "Decrement",
			Function: func(uc UpdateContext) (any, error) {
				self, ok := uc.SelfModel.(CounterModel)
				if !ok {
					return nil, errors.New("ERROR: Assertion failed")
				}
				return self.Data["count"].(int) - 1, nil
			},
		},
		"z": {
			KeyHint: "z",
			Legend:  "Reset to 0",
			Function: func(uc UpdateContext) (any, error) {
				return 0, nil
			},
		},
	},
	ReRenderPolicy: ReRenderOnChange,
	Focusable:      FocusTrue,
	Dependents:     []int{},
	UpdateFunc:     nil,
}

type CounterModel struct {
	Width, Height, Originx, Originy int
	Controls                        Command
	Data                            Data
	ReRenderPolicy                  ReRenderPolicy
	Focusable                       Focusable
	Dependents                      []int
	UpdateFunc                      UpdateFunc
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
		return u.SelfModel, false
	}
	if self.UpdateFunc != nil {
		return self.UpdateFunc(u)
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

	if userModel.UpdateFunc == nil {
		userModel.UpdateFunc = defaultBoxModel.UpdateFunc
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
