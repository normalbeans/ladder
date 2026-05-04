package component

type LState struct {
	CompModels                      map[int]Model
	Focus                           int
	WIDTH, HEIGHT, CURSORX, CURSORY int
}

type ComponentFunction struct {
	Legend   string
	Function func(LState) LState
}

type Command struct {
	Actions map[string]ComponentFunction
}

type Model interface{}

type Component interface {
	SetID(int) 
	Render(LState)
	Controls() Command
}

