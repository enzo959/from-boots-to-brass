package auth

import (
	"fmt"
)

func Marchand() {
	var choiceOption string
	for {
		fmt.Println("Vous etes chez le marchand\n1 - Voir les Objet en vente\n2 - Allé au Menu de jeu\n3 - Allé au Menu Principal")
		fmt.Scan(&choiceOption)

		switch choiceOption {
		case "1", "Voir les Objet en Vente":
			w1, w2, w3, e1, e2, e3, l1, l2 := Update()
			fmt.Print(w1, w2, w3, e1, e2, e3, l1, l2)
			BackMarchand()
		case "2", "Allé au Menu de jeu":
			BackMenuJeu()
		case "3", "Allé au Menu Principal":
			BackMenuPrincipal()
		default:
			Invalid()
		}
	}
}
