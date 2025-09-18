package auth

import (
	"fmt"
)

func Fight(t3, t4 Text) {
	var choiceOption string

	for {
		fmt.Println("Menu de jeu:\n1 - Rêgles\n2 - Crée son personage\n3 - Combat\n4 - Learn\n5 - Marchand\n6 - Allé au Menu Principal")
		fmt.Scan(&choiceOption)

		switch choiceOption {
		case "1", "Rêgles":
			fmt.Print(t3)
			Back()
		case "2", "Crée son personage":
			InitCharacter()
		case "3", "Combat":
			t5 := Book3()
			Battle(t5)
		case "4", "Learn":
			l1, l2 := Learn()
			fmt.Println(l1, l2)
			Back()
		case "5", "Marchand":
			w1, w2, w3, e1, e2, e3, l1, l2 := Update()
			fmt.Println(w1, w2, w3, e1, e2, e3, l1, l2)
			Back()
		case "6", "Allé au Menu Principal":
			t6 := Book4()
			fmt.Println(t6)
			BackMenuPrincipal()
		default:
			Invalid()
		}
	}
}

func Battle(t5 Text) {
	var choiceOption string

	for {
		fmt.Println("Action de jeu:\n1 - Attaque\n2 - Fuite")
		fmt.Scan(&choiceOption)

		switch choiceOption {
		case "1", "Attaque":
			fmt.Print(t5)
		case "2", "Fuite":
			BackMenuJeu()
		}
	}
}
