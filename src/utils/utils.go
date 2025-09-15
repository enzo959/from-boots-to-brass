package utils

import (
	"fmt"
)

func Clear() { // func clean terminal
	fmt.Print("\033[H\033[2J")
}
