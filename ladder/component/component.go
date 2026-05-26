package component

import (
	"context"
)

// focusable
type Focusable int

const (
	FocusUnset = iota
	FocusTrue
	FocusFalse
)

// render policies
type ReRenderPolicy int

const (
	ReRenderUnset = iota
	ReRenderOnChange
	ReRenderNever
	ReRenderAlways
)

type RenderContext struct {
	SelfModel  Model
	FocusModel Model
}

type UpdateContext struct {
	SelfModel  Model
	KeyBinding string
	Read       func(id int) Model
	Data       any
}

// Component data model interface.
type Model interface {
	GetControls() Command
	GetReRenderPolicy() ReRenderPolicy
	IsFocusable() Focusable
	GetDependents() []int // avoid self dependency
	GetBackgroundFunc() []BackgroundFunc
	Update(UpdateContext) (Model, bool)
}

// Funtion struct for component behaviour
type ComponentFunction struct {
	KeyHint  string
	Legend   string
	Function func(UpdateContext) (any, error)
}

type Command map[string]ComponentFunction

type Data map[string]any

// Component ( not model, but renderer that uses the model)
type Component interface {
	SetID(int)
	GetID() int
	Render(RenderContext)
	Init(Model) Model
	// DataModel(...) Model - not part of inteface but highly recommended for registering
}

// background running functions
type BackgroundMsg struct {
	ID   int
	Data any
}

type BackgroundFunc func(ctx context.Context, id int, c chan<- BackgroundMsg)

type UpdateFunc func(u UpdateContext) (Model, bool)
