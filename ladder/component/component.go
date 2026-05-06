package component

type LState struct {
	CompModels                      map[int]Model
	Changed                         map[int]bool
	Focus                           int
	WIDTH, HEIGHT, CURSORX, CURSORY int
}

type ComponentFunction struct {
	KeyHint  string
	Legend   string
	Function func(LState) LState
}

type Command struct {
	Actions map[string]ComponentFunction
}

type Model interface {
	GetControls() Command
}

type Component interface {
	SetID(int)
	GetID() int
	Render(LState)
	// Controls() Command
}
