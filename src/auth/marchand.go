package auth

import (
	"fmt"
)

func Marchand(perso *Character) {
	var choiceOption string
	for {
		fmt.Println("Vous etes chez le marchand\n1 - Voir les Objet en vente\n2 - Allé au Menu de jeu")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Voir les Objet en Vente":
			BuyItem(perso)
			Back()
		case "2", "Allé au Menu de jeu":
			return
		default:
			Invalid()
		}
	}
}

func BuyItem(perso *Character) {

	w1, w2, w3, e1, e2, e3, l1, l2 := Update() // Récupère les objets disponibles

	items := []StoreItem{
		{w1.Name, w1.Price}, {w2.Name, w2.Price}, {w3.Name, w3.Price}, {e1.Name, e1.Price}, {e2.Name, e2.Price}, {e3.Name, e3.Price}, {l1.Name, l1.Price}, {l2.Name, l2.Price}, // Crée une slice qui contient nom + prix
	}

	fmt.Println("Objets disponibles :") // Affiche chaque objet avec index et prix
	for i, it := range items {
		fmt.Printf("%d - %s (Prix : %d)\n", i+1, it.Name, it.Price)
	}

	fmt.Printf("\nVous avez %d crédit.\n", perso.Money)
	fmt.Print("Entrez le numéro de l’objet à acheter, si vous ne voulez rien entrez '0' : ")

	var num int
	fmt.Scanln(&num)

	if num > 0 && num <= len(items) {
		chosen := items[num-1]

		if perso.Money < chosen.Price { // Vérifie l’argent
			fmt.Printf("Vous ne possèdez pas assé de crédit pour acheter cette objet (°-°)\n")
			return
		}

		perso.Money -= chosen.Price // Déduit l’argent et ajoute à l’inventaire
		perso.Bag = append(perso.Bag, chosen.Name)

		fmt.Printf("%s acheté pour %d crédit. Vous possèdez à présent %dcrédit (^-^)/.\n",
			chosen.Name, chosen.Price, perso.Money)

		AccessInventory(*perso) // Affiche l’inventaire
	}
}
