package component

import (
	"fmt"
	"sort"
	"strings"
)

type LegendModel struct {
	Width, Height, Originx, Originy int
	Data                            Data
	Controls                        Command
	RenderPolicy                    RenderPolicy
}

func (l LegendModel) GetControls() Command {
	return l.Controls
}

func (l LegendModel) GetRenderControls() RenderPolicy {
	return l.RenderPolicy
}

func (l LegendModel) GetBackgroundFunc() []BackgroundFunc {
	return nil
}

func (l LegendModel) AsyncUpdate(data any) (Model, bool) {
	return l, false
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

// Interface implementation

func (l *Legend) SetID(id int) {
	l.id = id
}

func (l Legend) GetID() int {
	return l.id
}

func (l *Legend) Render(state LState) {
	c := state.CompModels[l.id].(LegendModel)

	lstring := generateLegendString(state.CompModels[state.Focus].GetControls())

	fmt.Printf("\x1b[%d;%dH\x1b[0K%s", c.Originy, c.Originx, lstring)
}

func (l *Legend) DataModel(width, height, ox, oy int, data Data, controls Command, rpolicy RenderPolicy) Model {
	return LegendModel{
		Width:        width,
		Height:       height,
		Originx:      ox,
		Originy:      oy,
		Data:         data,
		Controls:     controls,
		RenderPolicy: RenderAlways,
	}
}
