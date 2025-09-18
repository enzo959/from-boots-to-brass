package auth

import (
	"fmt"
	"os"
)

func Menu(t1, t2 Text) {
	var choiceOption string
	for {
		fmt.Println("Menu:\n1 - Prologue\n2 - Rules\n3 - Game menu\nTo quit the game, press the ‘Q’ key. ;)")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Prologue":
			fmt.Print(t1)
			Back()
		case "2", "Rules":
			fmt.Print(t2)
			Back()
		case "3", "Game menu":
			t3, t4 := Book2()
			Fight(t3, t4)
		case "q", "Q": // Exit
			fmt.Println("Good By !")
			os.Exit(0)
		default:
			Invalid()
		}
	}
}
