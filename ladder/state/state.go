package state

import "sync"

// state update and redraw

type GlobalState struct {
	mode            int // can be made simpler, needs more exploration
	activePage      int
	activeComponent int
}

func (g GlobalState) GetActivePage() int {
	return g.activePage
}

func (g GlobalState) GetActiveComponent() int {
	return g.activeComponent
}

func (g GlobalState) GetCurrentMode() int {
	return g.mode
}

func (g *GlobalState) SetScreenMode() {
	g.mode = 0
}

func (g *GlobalState) SetPageMode() {
	g.mode = 1
}

func (g *GlobalState) SetComponentMode() {
	g.mode = 2
}

func InitGlobalState() {
	globalstate = GlobalState{
		mode:            1,
		activePage:      0,
		activeComponent: 0,
	}
}

var (
	statelock   sync.RWMutex
	globalstate GlobalState
)

func UpdateState(newstate GlobalState) {
	statelock.Lock()
	globalstate = newstate
	statelock.Unlock()
}

func UpdateMode(mode int) {
	statelock.Lock()
	globalstate.mode = mode
	statelock.Unlock()
}

func UpdateActivePage(page int) {
	statelock.Lock()
	globalstate.activePage = page
	statelock.Unlock()
}

func UpdateActiveComponent(cmp int) {
	statelock.Lock()
	globalstate.activeComponent = cmp
	statelock.Unlock()
}

func GetState() GlobalState {
	statelock.RLock()
	defer statelock.RUnlock()
	return globalstate
}
