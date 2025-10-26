package textbox

import "fmt"

type TextBox struct {
	text string
}

func (t *TextBox) Init(text string) {
	t.text = text
}

func (t *TextBox) Redraw() {
	fmt.Print("\n")
	fmt.Print(t.text, "\n")
}

// can probably be merged, temp for now
func (t *TextBox) CalculateHeight() int {
	return 2
}
