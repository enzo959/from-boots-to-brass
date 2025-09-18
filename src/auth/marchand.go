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
	w1, w2, w3, e1, e2, e3, l1, l2 := Update()

	// Construis une slice d'Item (nom, type, puissance, prix)
	store := []Item{
		{w1.Name, "weapon", w1.PvAttack, w1.Price},
		{w2.Name, "weapon", w2.PvAttack, w2.Price},
		{w3.Name, "weapon", w3.PvAttack, w3.Price},
		{e1.Name, "heal", e1.PvSup, e1.Price},
		{e2.Name, "heal", e2.PvSup, e2.Price},
		{e3.Name, "heal", e3.PvSup, e3.Price},
		{l1.Name, "heal", l1.PvAttack, l1.Price},
		{l2.Name, "weapon", l2.PvAttack, l2.Price}, // adapte si l2 est loot/weapon
	}

	fmt.Println("Objets disponibles :")
	for i, it := range store {
		fmt.Printf("%d - %s (%s) — Prix : %d\n", i+1, it.Name, it.Kind, it.Price)
	}

	fmt.Printf("\nVous avez %d pièces.\n", perso.Money)
	fmt.Print("Entrez le numéro de l’objet à acheter, si vous ne voulez rien entrez 0: ")
	var num int
	fmt.Scanln(&num)

	if num > 0 && num <= len(store) {
		chosen := store[num-1]

		if perso.Money < chosen.Price {
			fmt.Printf("Pas assez d'argent pour acheter %s (prix %d).\n", chosen.Name, chosen.Price)
			return
		}

		// retire l'argent et ajoute l'objet (Item) dans l'inventaire
		perso.Money -= chosen.Price
		perso.Bag = append(perso.Bag, chosen)

		fmt.Printf("%s a été ajouté dans votre inventaire ;)\n", chosen.Name)
		AccessInventory(perso)
	}
}
