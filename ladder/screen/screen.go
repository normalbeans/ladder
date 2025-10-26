package screen

import (
	"github.com/nn-advith/ladder/ladder/page"
	"github.com/nn-advith/ladder/ladder/state"
)

type Screen struct {
	PageStack []page.Page
}

func (s *Screen) Init() {
	// new page here
}

func (s *Screen) RenderScreen(currentstate state.GlobalState) {
	s.PageStack[currentstate.GetActivePage()].RenderPage(currentstate)
}
