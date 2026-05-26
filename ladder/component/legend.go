package component

import (
	"fmt"
	"maps"
	"sort"
	"strings"
)

var defaultLegendModel = LegendModel{
	Width:          100,
	Height:         30,
	Originx:        1,
	Originy:        1,
	Data:           nil,
	Controls:       nil,
	ReRenderPolicy: ReRenderAlways,
	Focusable:      FocusFalse,
	Dependents:     nil,
}

type LegendModel struct {
	Width, Height, Originx, Originy int
	Data                            Data
	Controls                        Command
	ReRenderPolicy                  ReRenderPolicy
	Focusable                       Focusable
	Dependents                      []int
}

func (l LegendModel) GetControls() Command {
	return l.Controls
}

func (l LegendModel) GetReRenderPolicy() ReRenderPolicy {
	return l.ReRenderPolicy
}

func (l LegendModel) GetBackgroundFunc() []BackgroundFunc {
	return nil
}

func (l LegendModel) Update(u UpdateContext) (Model, bool) {
	// data can be anything
	return u.SelfModel, false
}

func (l LegendModel) IsFocusable() Focusable {
	return l.Focusable
}

func (l LegendModel) GetDependents() []int {
	return l.Dependents
}

func (l LegendModel) usingDefault(m Model) Model {
	userModel := m.(LegendModel)

	if userModel.Width == 0 {
		userModel.Width = defaultLegendModel.Width
	}
	if userModel.Height == 0 {
		userModel.Height = defaultLegendModel.Height
	}
	if userModel.Originx == 0 {
		userModel.Originx = defaultLegendModel.Originx
	}
	if userModel.Originy == 0 {
		userModel.Originy = defaultLegendModel.Originy
	}

	if userModel.Data == nil {
		userModel.Data = defaultLegendModel.Data
	} else {
		tempData := make(Data)
		maps.Copy(tempData, defaultLegendModel.Data)
		maps.Copy(tempData, userModel.Data)
		userModel.Data = tempData
	}

	if userModel.Controls == nil {
		userModel.Controls = defaultLegendModel.Controls
	} else {
		tempControls := make(Command)
		maps.Copy(tempControls, defaultLegendModel.Controls)
		maps.Copy(tempControls, userModel.Controls)
		userModel.Controls = tempControls
	}

	if userModel.ReRenderPolicy == ReRenderUnset {
		userModel.ReRenderPolicy = defaultLegendModel.ReRenderPolicy
	}

	if userModel.Focusable == FocusUnset {
		userModel.Focusable = defaultLegendModel.Focusable
	}

	if userModel.Dependents == nil {
		userModel.Dependents = defaultLegendModel.Dependents
	}

	return userModel
}

type Legend struct {
	id int
}

func generateLegendString(c Command) string {
	var focusindex strings.Builder

	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] <= keys[j]
	})

	for i := range keys {
		fmt.Fprintf(&focusindex, " %s - %s |", c[keys[i]].KeyHint, c[keys[i]].Legend)
	}
	return focusindex.String()
}

// // Interface implementation

func (l *Legend) SetID(id int) {
	l.id = id
}

func (l Legend) GetID() int {
	return l.id
}

func (l *Legend) Init(m Model) Model {

	newBoxModel := defaultLegendModel.usingDefault(m)

	return newBoxModel
}

func (l *Legend) Render(r RenderContext) {
	c := r.SelfModel.(LegendModel)
	focusModel := r.FocusModel
	lstring := generateLegendString(focusModel.GetControls())

	fmt.Printf("\x1b[%d;%dH\x1b[0K%s", c.Originy, c.Originx, lstring)
}
