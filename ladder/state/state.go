package state

import "sync"

// state update and redraw

type GlobalState struct {
	modeStack       []int
	activePage      int
	activeComponent int
}

func (g GlobalState) GetActivePage() int {
	return g.activePage
}

func (g GlobalState) GetActiveComponent() int {
	return g.activeComponent
}

func InitGlobalState() {
	globalstate = GlobalState{
		modeStack:       []int{0, 1, 2},
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

func GetState() GlobalState {
	statelock.RLock()
	defer statelock.RUnlock()
	return globalstate
}
