package textbox

import "fmt"

type TextBox struct {
	id            int
	text          string
	currentHeight int
}

func (t *TextBox) Init(id int, text string) {
	t.id = id
	t.text = text
}

func (t TextBox) Render() {
	fmt.Print(t.text)
	fmt.Print("\r\n")
}

// can probably be merged, temp for now
func (t TextBox) GetHeight() int {
	return t.currentHeight
}
