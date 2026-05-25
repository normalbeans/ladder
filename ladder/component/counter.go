package component

// import (
// 	"fmt"
// 	"maps"
// )

// type CounterModel struct {
// 	Width, Height, Originx, Originy int
// 	Controls                        Command
// 	Data                            Data
// 	RenderPolicy                    RenderPolicy
// }

// func (c CounterModel) GetControls() Command {
// 	return c.Controls
// }

// func (c CounterModel) GetRenderControls() RenderPolicy {
// 	return c.RenderPolicy
// }

// func (c CounterModel) GetBackgroundFunc() []BackgroundFunc {
// 	return nil
// }

// func (c CounterModel) AsyncUpdate(data any) (Model, bool) {
// 	return c, false
// }

// type Counter struct {
// 	id int
// }

// // Implement interface
// func (c Counter) GetID() int {
// 	return c.id
// }

// func (c *Counter) SetID(newID int) {
// 	c.id = newID
// }

// func (c *Counter) Render(state LState) {
// 	m := state.CompModels[c.id].(CounterModel)
// 	fmt.Printf("\x1b[%d;%dH", m.Originy, m.Originx)
// 	for i := 0; i < m.Height; i++ {
// 		if i == m.Height/2 {
// 			fmt.Printf("\x1b[%d;%dH\x1b[0K\tCount: %d", m.Originy+i, 1, m.Data["count"])
// 		} else {
// 			fmt.Printf("\x1b[%d;%dH", m.Originy+i, 1)
// 		}
// 	}
// }

// // not part of interface -> change this

// func (c *Counter) DataModel(width, height, ox, oy int, data Data, controls Command) Model {

// 	defaultData := Data{
// 		"count": 0,
// 	}

// 	defaultControls := Command{
// 		"+": {
// 			KeyHint: "+",
// 			Legend:  "Increment",
// 			Function: func(state LState) LState {
// 				cstate := state.CompModels[c.GetID()].(CounterModel)
// 				cstate.Data["count"]++
// 				state.CompModels[c.GetID()] = cstate
// 				return state
// 			},
// 		},
// 		"-": {
// 			KeyHint: "-",
// 			Legend:  "Decrement",
// 			Function: func(state LState) LState {
// 				cstate := state.CompModels[c.GetID()].(CounterModel)
// 				cstate.Data["count"]--
// 				state.CompModels[c.GetID()] = cstate
// 				return state
// 			},
// 		},
// 		"z": {
// 			KeyHint: "z",
// 			Legend:  "Reset to 0",
// 			Function: func(state LState) LState {
// 				cstate := state.CompModels[c.GetID()].(CounterModel)
// 				cstate.Data["count"] = 0
// 				state.CompModels[c.GetID()] = cstate
// 				return state
// 			},
// 		},
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
