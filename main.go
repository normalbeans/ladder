package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Message any

type Increment struct{}
type Decrement struct{}
type lMsg struct {
	id  int
	msg any
}

type Component interface {
	Init(Message)
	Update(Message) bool
	View() string
}

type Counter struct {
	count int
}

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
		if m == "r" {
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
	sb.WriteString("\r\nCounter:")
	fmt.Fprintf(&sb, "\r\nCount: %d\r\n=========\r\n", c.count)
	return sb.String()
}

type Ladder struct {
	// channels
	components    []Component
	msgChannel    chan Message
	focus         int
	terminalState *term.State
}

func (l *Ladder) Init(c []Component) {
	l.components = c
	l.focus = 0
	l.msgChannel = make(chan Message)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	l.terminalState = oldState

	go func() {
		b := make([]byte, 1)
		for {

			os.Stdin.Read(b)
			input := string(b)
			l.msgChannel <- input
		}
	}()
}

// func (l *Ladder) AddComponent(c Component) {
// 	l.components = append(l.components, c)
// }

func (l *Ladder) Update(msg Message) {

	switch m := msg.(type) {
	case string:
		switch m {
		case "q", "\x03":
			term.Restore(int(os.Stdin.Fd()), l.terminalState)
			os.Exit(0)
		case "j":
			l.focus = (l.focus - 1) % len(l.components)
		case "k":
			l.focus = (l.focus + 1) % len(l.components)
		case "+":
			l.components[l.focus].Update(Increment{})
		case "-":
			l.components[l.focus].Update(Decrement{})
		default:
			l.components[l.focus].Update(m)
		}
	}
}

func (l *Ladder) View() {
	// get component strings
	var sb strings.Builder
	sb.WriteString("\033[H\033[2J")
	for i := range l.components {
		sb.WriteString(l.components[i].View())
		sb.WriteString("\r\n")
	}
	fmt.Print(sb.String())

}

func (l *Ladder) Render() {
	l.View()

	for m := range l.msgChannel {
		l.Update(m)
		l.View()
	}
}

// ladder run must initialise the models,

func main() {

	// msgChannel := make(chan Message)
	c1 := &Counter{}
	c1.Init(10)
	c2 := &Counter{}
	c2.Init(20)

	l := &Ladder{}
	l.Init([]Component{c1, c2})
	l.Render()

}
