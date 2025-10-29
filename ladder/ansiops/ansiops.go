package ansiops

import "fmt"

func ClearScreen() {
	fmt.Print("\033[2J\033[3J\033[H")
}
