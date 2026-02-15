package page

import (
	"github.com/nn-advith/ladder/ladder/component"
	"github.com/nn-advith/ladder/ladder/state"
)

type Page struct {
	ComponentStack []component.Component
}

func (p *Page) RenderPage(currentstate state.GlobalState) {
	// fmt.Print("\033[2J\033[3J\033[H")
	for i := range p.ComponentStack {
		p.ComponentStack[i].Render()
	}
}
