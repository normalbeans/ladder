package state

import "github.com/normalbeans/ladder/ladder/component"

// Global state owned by ladder
type LState struct {
	CompModels                      map[int]component.Model
	Changed                         map[int]bool
	Focus                           int
	WIDTH, HEIGHT, CURSORX, CURSORY int // CURSORX AND CURSORY are not used.
}
