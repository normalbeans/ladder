package ladder

import (
	"os"

	"github.com/nn-advith/ladder/ladder/component"
	"github.com/nn-advith/ladder/ladder/component/textbox"
	"github.com/nn-advith/ladder/ladder/page"
	"github.com/nn-advith/ladder/ladder/screen"
	"github.com/nn-advith/ladder/ladder/state"
	"golang.org/x/term"
)

// TODO:
// create channels for redraw and quit; redraw channel is of type GlobalState.

var renderchannel = CreateRenderChannel()
var inputchannel = CreateInputChannel()

func CreateRenderChannel() chan state.GlobalState {
	return make(chan state.GlobalState, 1)
}

func CreateInputChannel() chan []byte {
	return make(chan []byte, 1)
}

type Ladder struct {
	Screen    screen.Screen
	TermState *term.State
}

func (l *Ladder) Render() { // listen to state here
	// fmt.Print("starting redner")
	cs := state.GetState()
	renderchannel <- cs
	for range renderchannel { // for range instead of for select because only one channel
		cs := state.GetState()
		// fmt.Println(cs)
		l.Screen.RenderScreen(cs)
	}
}

func InitLadder() *Ladder {
	// initialise a ladder object containing one screen, and data component
	// put terminal to raw mode and do something

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// init state
	state.InitGlobalState()

	newtextbox := textbox.TextBox{}
	newtextbox.Init("Hello world")

	newScreen := screen.Screen{
		PageStack: []page.Page{
			{
				ComponentStack: []component.Component{
					&newtextbox,
				},
			},
		},
	}
	return &Ladder{
		TermState: oldState,
		Screen:    newScreen,
	}
}

func InputListener() {
	buf := make([]byte, 3)
	for {
		for i := range buf {
			buf[i] = 0
		}
		_, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		switch buf[0] {
		case 13, 10:
			// enter/return
			cs := state.GetState()
			renderchannel <- cs
		case 'q', 'Q':
			return
		}
	}
}
