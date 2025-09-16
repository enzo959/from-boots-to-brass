package auth

import (
	"fmt"
	"projet-red_from-boots-to-brass/src/utils"
	"time"
)

type Text struct {
	text string
}

func Book() (Text, Text, Inventory) {
	t1 := Text{" Prlogue: \n Dans ce monde, la gloire n’est pas donnée. Elle se prend, balle après balle. \n Je ne suis qu’un soldat parmi des milliers : un matricule, une arme, un ordre. \n On ne choisit pas où l’on commence. On choisit jusqu’où on ira. \n Les tirs résonnent au loin, les hélicoptères déchirent le ciel. Aujourd’hui, ce n’est plus qu’une mission : c’est le premier pas. \n Le champ de bataille forge les hommes… et je compte bien gravir chaque échelon, quel qu’en soit le prix."}
	t2 := Text{"Rêgle: \n test Regles"}
	t3 := Inventory{"Inventaire:\n", []string{"text", "text"}}
	return t1, t2, t3
}

func Menu (t1 Text, t2 Text, t3 Inventory) {
	var choiceOption string
	var option // choice type 

	for {
		fmt.Println("Menu:\n1 - Prologue\n2 - Rules\n3 - Game\nPour quitter le jeu lisez le prologue ;)")
		fmt.Scan(&choiceOption)

		switch choiceOption {
			case "1", "Prologue":
				option = t1
			case "2", "Rules":
				option = t2
			case "3", "Game":
				option = t3 // contraint type
			default:
				utils.Invalid()	
		}
		utils.Clear()
		fmt.Printf("\n%s\n", option.text)
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
