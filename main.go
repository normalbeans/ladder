package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var CURSORX int = 1
var CURSORY int = 1

func printAtPos(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
	fmt.Print("AAA")
}

func printRectangle(ox, oy, w, h int) {
	// prints a rectangle of width w and height h starting from ox and oy.
	fmt.Printf("\033[%d;%dH", ox, oy)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf("\033[%d;%dH\u2580", ox+i, oy+j)
		}
	}
}

func printRectangleAtCurrent(w, h int) {
	for j := 0; j < h; j++ {
		fmt.Printf("\033[%d;%dH%s", j+1, 1, strings.Repeat("\u2580 ", w))
	}
}

func displayCoord() {
	fmt.Printf("\033[%d;%dH\033[0KCord: (%d, %d)\t| Ctrl+C - Quit", 15, 1, CURSORX, CURSORY)
}

func snoopy(quit chan byte, w, h int) {
	// just listen keystrokes and print them back
	// topblock := true
	// snooper := bufio.NewReader(os.Stdin)
	for {

		fmt.Printf("\033[%d;%dH", CURSORY, CURSORX)

		buf := make([]byte, 8)
		n, err := os.Stdin.Read(buf)
		data := buf[:n]
		if err != nil {
			panic(err)
		}
		if data[0] == 0x03 {
			close(quit)
		} else if data[0] == 27 && data[1] == 91 {
			switch data[2] {
			case 65:
				CURSORY = max(CURSORY-1, 1)
				// topblock = !topblock
			case 66:
				CURSORY = min(CURSORY+1, h)
				// topblock = !topblock
			case 67:
				CURSORX = min(CURSORX+1, w)
			case 68:
				CURSORX = max(CURSORX-1, 1)
			}
			displayCoord()
		} else {
			if CURSORX < w {
				CURSORX = min(CURSORX+1, w)
				fmt.Print(string(data))
				displayCoord()
			}
		}

	}
}

func main() {

	WIDTH, HEIGHT := 30, 14

	originalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), originalState)
	quit := make(chan byte, 1)
	fmt.Print("\033[H\033[J")
	go snoopy(quit, WIDTH, HEIGHT)
	displayCoord()
	<-quit
	fmt.Print("\033[H\033[J\033[?25h")
}
