package auth

import (
	"fmt"
)

type Item struct {
	Name  string // nom affiché
	Kind  string // "weapon" ou "heal"
	Power int    // puissance (dégâts ou points de vie rendus)
	Price int
}

type Character struct { // create class character
	Name  string
	Class string
	Level int
	PvMax int
	PvAct int
	Money int
	Bag   []Item
}

type Inventory struct {
	Text string
	Inv  []string
}

func InitCharacter() Character {
	var name, class string

	fmt.Print("Entrez le nom de votre soldat: ")
	fmt.Scanln(&name)

	fmt.Print("Entrez la nationalité (classe) de votre soldat: ")
	fmt.Scanln(&class)

	return Character{
		Name:  name,
		Class: class,
		Level: 1,
		PvMax: 200,
		PvAct: 100,
		Money: 50,
		Bag:   []Item{},
	}
}

func AccessInventory(perso *Character) {
	if len(perso.Bag) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, it := range perso.Bag {
		fmt.Printf("%d - %s (%s, puissance %d)\n", i+1, it.Name, it.Kind, it.Power)
	}
}
