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

// Component data model interface.
type Model interface {
	GetControls() Command
	GetReRenderPolicy() ReRenderPolicy
	IsFocusable() Focusable
	GetDependents() []int // avoid self dependency
	GetBackgroundFunc() []BackgroundFunc
	Update(Model, any) (Model, bool)
}

// Funtion struct for component behaviour
type ComponentFunction struct {
	KeyHint  string
	Legend   string
	Function func(data any) any
}

type Command map[string]ComponentFunction

type Data map[string]any

// Component ( not model, but renderer that uses the model)
type Component interface {
	SetID(int)
	GetID() int
	Render(Model)
	Init(Model) Model
	// DataModel(...) Model - not part of inteface but highly recommended for registering
}

// background running functions
type BackgroundMsg struct {
	ID   int
	Data any
}

type BackgroundFunc func(ctx context.Context, id int, c chan<- BackgroundMsg)
