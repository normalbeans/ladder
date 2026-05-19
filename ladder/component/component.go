package component

import "context"

// Global state owned by ladder
type LState struct {
	CompModels                      map[int]Model
	Changed                         map[int]bool
	Focus                           int
	WIDTH, HEIGHT, CURSORX, CURSORY int // CURSORX AND CURSORY are not used.
}

// render policies
type RenderPolicy int

const (
	RenderAlways = iota
	RenderNever
)

// Component data model interface.
type Model interface {
	GetControls() Command
	GetRenderControls() RenderPolicy
	GetBackgroundFunc() []BackgroundFunc
	AsyncUpdate(data any) (Model, bool)
}

// Funtion struct for component behaviour
type ComponentFunction struct {
	KeyHint  string
	Legend   string
	Function func(LState) LState
}

type Command map[string]ComponentFunction

type Data map[string]int

// Component ( not model, but renderer that uses the model)
type Component interface {
	SetID(int)
	GetID() int
	Render(LState)
	// DataModel(...) Model - not part of inteface but highly recommended for registering
}

// background running functions
type BackgroundMsg struct {
	ID   int
	Data any
}

type BackgroundFunc func(ctx context.Context, id int, c chan<- BackgroundMsg)
