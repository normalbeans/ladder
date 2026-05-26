package component

import (
	"context"
	"fmt"
	"maps"
	"math/rand/v2"
	"strings"
	"time"
)

// this is to help with Init command.
var defaultBoxModel = BoxModel{
	Width:   100,
	Height:  5,
	Originx: 1,
	Originy: 1,
	Data: Data{
		"color": getRandomColor(),
	},
	Controls: Command{
		"c": {
			KeyHint: "c",
			Legend:  "Randomise color",
			// Function: func(data any) any {
			// 	return getRandomColor()
			// },
		},
	},
	ReRenderPolicy: ReRenderOnChange,
	Focusable:      FocusTrue,
	Dependents:     []int{},
}

type BoxModel struct {
	Width, Height    int
	Originx, Originy int
	Data             Data
	Controls         Command
	ReRenderPolicy   ReRenderPolicy
	Focusable        Focusable
	Dependents       []int
}

type Box struct {
	id int
}

// sample background func
func setRandomColorOnTimer(ctx context.Context, id int, ch chan<- BackgroundMsg) {
	timer := time.NewTicker(3 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			ch <- BackgroundMsg{ID: id, Data: getRandomColor()}
		case <-ctx.Done():
			return
		}
	}
}

func getRandomColor() int {
	return 31 + rand.IntN(5)
}

func (b BoxModel) GetControls() Command {
	return b.Controls
}

func (b BoxModel) GetReRenderPolicy() ReRenderPolicy {
	return b.ReRenderPolicy
}

func (b BoxModel) IsFocusable() Focusable {
	return b.Focusable
}

func (b BoxModel) GetDependents() []int {
	return b.Dependents
}

func (b BoxModel) GetBackgroundFunc() []BackgroundFunc {
	return []BackgroundFunc{setRandomColorOnTimer}
}

func (b BoxModel) Update(u UpdateContext) (Model, bool) {
	self, ok := u.SelfModel.(BoxModel)
	if !ok {
		return b, false
	}
	if u.KeyBinding != "" {

		switch u.KeyBinding {
		case "c":
			self.Data["color"] = getRandomColor()
		default:
			return self, false
		}
	}
	if u.Data != nil {
		c, ok := u.Data.(int)
		if !ok {
			return self, false
		}
		self.Data["color"] = c
	}
	return self, true
}

// helper to merge defaults to model passed from user.
// TODO: can be made a common util since other components prolly also will retain same fields
func (b BoxModel) usingDefault(m Model) Model {
	userModel := m.(BoxModel)

	if userModel.Width == 0 {
		userModel.Width = defaultBoxModel.Width
	}
	if userModel.Height == 0 {
		userModel.Height = defaultBoxModel.Height
	}
	if userModel.Originx == 0 {
		userModel.Originx = defaultBoxModel.Originx
	}
	if userModel.Originy == 0 {
		userModel.Originy = defaultBoxModel.Originy
	}

	if userModel.Data == nil {
		userModel.Data = defaultBoxModel.Data
	} else {
		tempData := make(Data)
		maps.Copy(tempData, defaultBoxModel.Data)
		maps.Copy(tempData, userModel.Data)
		userModel.Data = tempData
	}

	if userModel.Controls == nil {
		userModel.Controls = defaultBoxModel.Controls
	} else {
		tempControls := make(Command)
		maps.Copy(tempControls, defaultBoxModel.Controls)
		maps.Copy(tempControls, userModel.Controls)
		userModel.Controls = tempControls
	}

	if userModel.ReRenderPolicy == ReRenderUnset {
		userModel.ReRenderPolicy = defaultBoxModel.ReRenderPolicy
	}

	if userModel.Focusable == FocusUnset {
		userModel.Focusable = defaultBoxModel.Focusable
	}

	if userModel.Dependents == nil {
		userModel.Dependents = defaultBoxModel.Dependents
	}

	return userModel
}

// Implement interface

func (b Box) GetID() int {
	return b.id
}

func (b *Box) SetID(id int) {
	b.id = id
}

func (b *Box) Render(r RenderContext) {
	c := r.SelfModel.(BoxModel)
	fmt.Printf("\x1b[%d;%dH", c.Originy, c.Originx)
	for j := 0; j < c.Height; j++ {
		switch j {
		case 0:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2580", c.Width-2)+"\u2588\x1b[0m", c.Data["color"], c.Originy+j, c.Originx)
		case c.Height - 1:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat("\u2584", c.Width-2)+"\u2588\x1b[0m", c.Data["color"], c.Originy+j, c.Originx)
		default:
			fmt.Printf("\x1b[%dm\x1b[%d;%dH\u2588"+strings.Repeat(" ", c.Width-2)+"\u2588\x1b[0m", c.Data["color"], c.Originy+j, c.Originx)
		}
	}
}

func (b *Box) Init(m Model) Model {

	newBoxModel := defaultBoxModel.usingDefault(m)

	return newBoxModel
}
