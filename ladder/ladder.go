package ladder

import (
	"fmt"
	"os"

	"github.com/nn-advith/ladder/ladder/ansiops"
	"github.com/nn-advith/ladder/ladder/compevent"
	"github.com/nn-advith/ladder/ladder/component"
	"github.com/nn-advith/ladder/ladder/component/lineinput"
	"github.com/nn-advith/ladder/ladder/component/textbox"
	"github.com/nn-advith/ladder/ladder/page"
	"github.com/nn-advith/ladder/ladder/screen"
	"github.com/nn-advith/ladder/ladder/state"
	"golang.org/x/term"
)

// TODO:
// create channels for redraw and quit; redraw channel is of type GlobalState.

var renderchannel = CreateRenderChannel() // handle state global
var compeventchannel = CreateCompEventChannel()
var inputdonechannel = CreateInputDoneChannel()

// var inputchannel = CreateInputChannel()

func CreateRenderChannel() chan state.GlobalState {
	return make(chan state.GlobalState, 1)
}

func CreateCompEventChannel() chan compevent.CompEvent {
	return make(chan compevent.CompEvent, 1)
}

func CreateInputDoneChannel() chan struct{} {
	return make(chan struct{})
}

// func CreateInputChannel() chan []byte {
// 	return make(chan []byte, 1)
// }

type Ladder struct {
	Screen    screen.Screen
	TermState *term.State
}

func (l *Ladder) Render() { // listen to state here
	// fmt.Print("starting redner")

	go func() {
		cs := state.GetState()
		renderchannel <- cs
		for range renderchannel { // for range instead of for select because only one channel
			cs := state.GetState()
			// fmt.Println(cs)
			l.Screen.RenderScreen(cs)
		}
	}()

	InputListener()
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

	// clear terminal
	ansiops.ClearScreen()

	newtextbox := textbox.TextBox{}
	newtextbox.Init(0, "Hello world")

	// commandpanel := commandpanel.CommandPanel{}
	// commandpanel.Init("UP: Move Up Down: Move Down ENTER: Select component")

	lineinput := lineinput.Lineinput{}
	lineinput.Init(1, "INPUT: ", compeventchannel, inputdonechannel)

	newScreen := screen.Screen{
		PageStack: []page.Page{
			{
				ComponentStack: []component.Component{
					// &commandpanel,
					&newtextbox,
					&lineinput,
				},
			},
		},
	}
	return &Ladder{
		TermState: oldState,
		Screen:    newScreen,
	}
}

// Input listener loop
func InputListener() {
	buf := make([]byte, 8)
	for {

		//; this is not being triggered immediately. check this. maybe due to os.std in
		// try moving read to separate routing and introduce YET another channel.
		select {
		case <-inputdonechannel:
			fmt.Print("\n\rinput done")
			state.UpdateMode(1)
			cs := state.GetState() // definitely improve this, maybe setup a different listener
			renderchannel <- cs
			continue
		default:
		}

		for i := range buf {
			buf[i] = 0
		}
		n, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		data := buf[:n]
		cs := state.GetState()
		cmode := cs.GetCurrentMode()
		fmt.Printf("\n\rCMODE: %d", cmode)

		if cmode == 1 {
			// page mode
			switch data[0] {
			case 13, 10:
				// enter/return
				state.UpdateMode(2)
				state.UpdateActiveComponent(1)
			case 'q', 'Q':
				return
			default:
				//pass
			}
		} else if cmode == 2 {
			//component mode
			ncevent := compevent.CompEvent{}
			ncevent.Id = cs.GetActiveComponent()
			valCopy := append([]byte(nil), data...)
			ncevent.Val = valCopy

			fmt.Printf("\n\rSENDING EVENT: %d, %v", ncevent.Id, ncevent.Val)
			compeventchannel <- ncevent
		}

	}
}
