package auth

import (
	"fmt"
)

func Marchand(perso *Character) {
	var choiceOption string
	for {
		fmt.Println("Vous etes chez le marchand\n1 - Voir les Objet en vente\n2 - Allé au Menu de jeu\n3 - Allé au Menu Principal")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Voir les Objet en Vente":
			BuyItem(perso)
			Back()
		case "2", "Allé au Menu de jeu":
			return
		case "3", "Allé au Menu Principal":
			return
		default:
			Invalid()
		}
	}
}

func BuyItem(perso *Character) { // Récupère les objets disponibles

	w1, w2, w3, e1, e2, e3, l1, l2 := Update() // Met tous les noms d’objets dans une slice

	items := []string{
		w1.Name, w2.Name, w3.Name,
		e1.Name, e2.Name, e3.Name,
		l1.Name, l2.Name,
	}

	fmt.Println("Objets disponibles :") // Affiche chaque objet avec un index
	for i, it := range items {
		fmt.Printf("%d - %s\n", i+1, it)
	}

	fmt.Print("Entrez le numéro de l’objet à acheter (0 pour annuler) : ") // Choix utilisateur
	var num int
	fmt.Scanln(&num)

	if num > 0 && num <= len(items) {
		chosen := items[num-1]
		perso.Bag = append(perso.Bag, chosen)
		fmt.Printf("%s a été ajouté à votre inventaire !\n", chosen)
		AccessInventory(*perso)
	}
}
