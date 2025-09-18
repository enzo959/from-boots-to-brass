package auth

import (
	"fmt"
)

func Marchand(perso *Character) {
	var choiceOption string
	for {
		fmt.Println("You are at the store\n1 - See Items for Sale\n2 - Go to Game Menu")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "See Items for Sale":
			BuyItem(perso)
			Back()
		case "2", "Go to Game Menu":
			return
		default:
			Invalid()
		}
	}
}

func BuyItem(perso *Character) {
	w1, w2, w3, e1, e2, e3, l1, l2 := Update()

	// Build an Item slice
	store := []Item{
		{w1.Name, "weapon", w1.PvAttack, w1.Price},
		{w2.Name, "weapon", w2.PvAttack, w2.Price},
		{w3.Name, "weapon", w3.PvAttack, w3.Price},
		{e1.Name, "heal", e1.PvSup, e1.Price},
		{e2.Name, "heal", e2.PvSup, e2.Price},
		{e3.Name, "heal", e3.PvSup, e3.Price},
		{l1.Name, "heal", l1.PvAttack, l1.Price},
		{l2.Name, "weapon", l2.PvAttack, l2.Price}, // fits between loot/weapon
	}

	fmt.Println("Available items :")
	for i, it := range store {
		fmt.Printf("%d - %s (%s) — Price : %d\n", i+1, it.Name, it.Kind, it.Price)
	}

	fmt.Printf("\nYou have %d credit.\n", perso.Money)
	fmt.Print("Enter the number of the item you want to purchase. If you do not want to purchase anything, enter 0: ")
	var num int
	fmt.Scanln(&num)

	if num > 0 && num <= len(store) {
		chosen := store[num-1]

		if perso.Money < chosen.Price {
			fmt.Printf("Not enough credit to buy %s (price %d).\n", chosen.Name, chosen.Price)
			return
		}

		// removes the credit and adds Item to the inventory
		perso.Money -= chosen.Price
		perso.Bag = append(perso.Bag, chosen)

		fmt.Printf("%s has been added to your inventory ;)\n", chosen.Name)
		AccessInventory(perso)
	}
}
