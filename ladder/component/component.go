package component

type Component interface {
	// Every component must be of this type
	GetHeight() int
	Render()
}
