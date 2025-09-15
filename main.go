package main

import (
	"projet-red_from-boots-to-brass/src/auth"
	"projet-red_from-boots-to-brass/src/utils"
	"time"
)

func main() {
	p1, p2, p3 := auth.InitCharacter()
	auth.CharacterCreation(p1, p2, p3) // start func InitCharacter
	time.Sleep(10 * time.Second)       // stop the programme during 5s
	utils.Clear()                      // start func Clear
}
