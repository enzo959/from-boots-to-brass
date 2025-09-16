package main

import (
	"projet-red_from-boots-to-brass/src/auth"
)

func main() {
	/*p1, p2, p3 := auth.DisplayCharacter()
	auth.CharacterCreation(p1, p2, p3) // start func InitCharacter
	time.Sleep(10 * time.Second)       // stop the programme during 5s
	utils.Clear()
	t1, t2, t3 := auth.Book() // start func Clear
	auth.Menu(t1, t2, t3)*/
	auth.InitCharacter()
}
