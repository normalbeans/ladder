package lineinput

import (
	"fmt"

	"github.com/nn-advith/ladder/ladder/compevent"
)

type Lineinput struct {
	id            int
	prompt        string
	input         string
	currentHeight int
}

func (l *Lineinput) Init(id int, prompt string, cechannel chan compevent.CompEvent, idchannel chan struct{}) {
	l.id = id
	l.prompt = prompt

	go l.Listener(cechannel, idchannel)
}

func (l *Lineinput) Listener(cechannel chan compevent.CompEvent, idchannel chan struct{}) {
	// fmt.Print("\n\r started listener for id", l.id)
	for ip := range cechannel {
		// fmt.Printf("\n\rgot event: %d, %v", ip.Id, ip.Val)
		if ip.Id != l.id {
			//for me
			continue
		}
		// handle the byte slice here
		done := l.HandleInput(ip)
		if done == 1 {
			fmt.Printf("\n\r%s\n\r", l.input)
			fmt.Print("\n\rsending done")
			idchannel <- struct{}{}
		}
	}
}

func (l *Lineinput) HandleInput(ce compevent.CompEvent) int {
	switch ce.Val[0] {
	case 13, 10:
		//done
		return 1
	case 127, 8:
		if len(l.input) > 0 {
			l.input = l.input[:len(l.input)-1]
			// if (len(input)+len("New task: "))%width == width-1 {
			// 	fmt.Printf("\033[A\033[%dC \033[%dD", width, 1)
			// 	fmt.Print("\033[1C")
			// } else {
			fmt.Print("\b \b")
			// }

		}
	case 27:
		l.input = ""
		return 1
	default:
		l.input = l.input + string(ce.Val)
	}
	return 0
}

func (l *Lineinput) UpdatePrompt(prompt string) {
	l.prompt = prompt
}

func (l *Lineinput) Render() {
	fmt.Printf("\r\n%s", l.prompt)
}

// can probably be merged, temp for now
func (l Lineinput) GetHeight() int {
	return l.currentHeight
}
