package commandpanel

import "fmt"

type CommandPanel struct {
	currentCommandMenu string
	currentHeight      int
}

func (c *CommandPanel) Init(text string) {
	c.currentCommandMenu = text
}

func (c CommandPanel) Render() {
	fmt.Print("\n\r", c.currentCommandMenu)
}

// can probably be merged, temp for now
func (c CommandPanel) GetHeight() int {
	return c.currentHeight
}
