package main

import (
	"projet-red_from-boots-to-brass/src/auth"
	"projet-red_from-boots-to-brass/src/utils"
	"time"
)

func main() {
	auth.InitCharacter()        // start func InitCharacter
	time.Sleep(5 * time.Second) //
	utils.Clear()               // start func Clear
}
