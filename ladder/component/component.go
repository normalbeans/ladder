package component

type Component interface {
	// Every component must be of this type
	CalculateHeight() int
	Redraw()
}
