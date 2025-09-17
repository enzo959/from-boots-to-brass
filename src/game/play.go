package game

import (
	"fmt"
	"projet-red_from-boots-to-brass/src/utils"
)

type Text struct {
	text string
}

func Fight(b1, b2 Text) {
	var choiceOption string
	var option Text

	for {
		fmt.Println("Menu:\n1 - Rêgles\n2 - Combat\n3 - Learn\n4 - Marchand\nPour quitter le jeu lisez le prologue ;)")
		fmt.Scan(&choiceOption)

		switch choiceOption {
		case "1", "Rêgles":
			fmt.Print(b1)
		case "2", "Combat":
			fmt.Println(b2)
		case "3", "Learn":
			l1, l2 := Learn()
			fmt.Println(l1, l2)
		case "4", "Marchand":
			w1, w2, w3, e1, e2, e3, l1, l2 := Update()
			fmt.Println(w1, w2, w3, e1, e2, e3, l1, l2)
		default:
			utils.Invalid()
		}
		utils.Clear()
		fmt.Printf("\n%s\n", option.text)
	}
}
