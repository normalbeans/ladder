package component

import "fmt"

type LegendModel struct {
	Width, Height, Originx, Originy int
	Separator                       string
	LegendString                    string
	Controls                        Command
}

func (l LegendModel) GetControls() Command {
	return l.Controls
}

type Legend struct {
	id int
}

func (l *Legend) GenerateLegendString() {

}

// Interface implementation

func (l *Legend) SetID(id int) {
	l.id = id
}

func (l *Legend) Render(state LState) {
	c := state.CompModels[l.id].(LegendModel)
	fmt.Printf("\x1b[%d;%dH", c.Originy, c.Originx)

}
