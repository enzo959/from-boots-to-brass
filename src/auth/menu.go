package auth

import (
	"fmt"
	"os"
)

func Menu(t1, t2 Text) {
	var choiceOption string
	for {
		fmt.Println("Menu:\n1 - Prologue\n2 - Rules\n3 - Menu de jeu\nPour quitter le jeu appuyez sur la touche 'Q' ;)")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Prologue":
			fmt.Print(t1)
			Back()
		case "2", "Rules":
			fmt.Print(t2)
			Back()
		case "3", "Menu de jeu":
			t3, t4 := Book2()
			Fight(t3, t4)
		case "q", "Q": // Exit
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			Invalid()
		}
	}
}

/*

//func displayinfo() {

//}

func Menu(t1, t2, t3 Text) Text {
	var choiceOption string
	var option Text

	for {
		fmt.Println("Menu:\n1 - Prologue\n2 - Rules\n3 - Game\nPour quitter le jeu lisez le prologue ;)")
		fmt.Scan(&choiceOption)

		switch choiceOption {
		case "1", "Prologue": // if user write "1" or "Prologue"
			option = t1 // option = t1
		case "2", "Rules":
			option = t2
		case "3", "Game":
			option = t3
		default:
			fmt.Println("Choisissez une option valide.") // error gestion
			time.Sleep(2 * time.Second)                  // reset after 2 sec
			utils.Clear()                                // clean terminal
			continue                                     // restart the code
		}
		utils.Clear()
		fmt.Printf("\n%s\n", option.text) // print

		var back string // var back for back to menu
		fmt.Println("\nAppuyez sur n'importe quelle lettres pour revenir au menu, ou 'q' pour quitter :")
		fmt.Scan(&back)
		if back == "q" {
			return option // exit definitiv
		} else {
			utils.Clear()
		}
	}
}*/
