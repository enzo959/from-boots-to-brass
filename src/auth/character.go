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

	fmt.Print("Enter your soldier's name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your soldier's nationality: ")
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
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("inventory :")
	for i, it := range perso.Bag {
		fmt.Printf("%d - %s (%s, power %d)\n", i+1, it.Name, it.Kind, it.Power)
	}
}
