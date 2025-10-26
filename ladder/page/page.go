package page

import (
	"github.com/nn-advith/ladder/ladder/component"
	"github.com/nn-advith/ladder/ladder/state"
)

type Page struct {
	ComponentStack []component.Component
}

func (p *Page) RenderPage(currentstate state.GlobalState) {
	for i := range p.ComponentStack {
		p.ComponentStack[i].Redraw()
	}
}
