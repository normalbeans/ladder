package main

import (
	"fmt"
	"strings"
)

type Message interface{}

type Model interface {
	Init(Message)
	Update(Message) bool
	View() string
}

type Counter struct {
	count int
}

type Increment struct{}
type Decrement struct{}

func (c *Counter) Init(m Message) {
	switch m := m.(type) {
	case int:
		c.count = m
	default:
		c.count = 0
	}
}

func (c *Counter) Update(m Message) bool {
	switch m := m.(type) {
	case Increment:
		c.count += 1
	case Decrement:
		c.count -= 1
	case string:
		if m == "RESET" {
			c.count = 0
		}
	default:
		//nothing
		return false
	}
	return true
}

func (c *Counter) View() string {
	var sb strings.Builder
	sb.WriteString("\nCounter:")
	fmt.Fprintf(&sb, "\nCount: %d\n=========", c.count)
	return sb.String()
}

// type Selector struct {
// 	choices []string
// 	selection int
// }

// func (s *Selector) Init(m Message) {
// 	switch m := m.(type) {
// 	case []string:
// 		s.choices = m
// 	default:
// 		s.choices = []string{}
// 	}
// 	s.selection = 0
// }
// func (s *Selector) Update(m Message) bool {

// }
// func (s *Selector) {

// }

func main() {
	// w, h, err := term.GetSize(int(os.Stdin.Fd()))
	// if err != nil {
	// 	fmt.Println("Errror during getting size", err)
	// }

	// oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	// if err != nil {
	// 	fmt.Println("Error during raw:", err)
	// }
	// defer term.Restore(int(os.Stderr.Fd()), oldState)

	// fmt.Print("\r\nHello world")
	// fmt.Printf("\r\nWidth: %d\r\nHeight: %d", w, h)
	// fmt.Printf("\r\nTerminal: %v\r\n", term.IsTerminal(int(os.Stdin.Fd())))
	c := Counter{}
	c.Init(nil)
	fmt.Println(c.View())
	for {
		rerender := false
		// wait for input
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "+":
			rerender = c.Update(Increment{})
		case "-":
			rerender = c.Update(Decrement{})
		case "r":
			rerender = c.Update("RESET")
		default:
			continue
		}

		if rerender {
			fmt.Println(c.View())
			rerender = false
		}
	}
}
