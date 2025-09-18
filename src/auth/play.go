package auth

import (
	"fmt"
	"math/rand"
)

func Fight(t3, t4 Text) {
	var choiceOption string
	var perso *Character
	for {
		fmt.Println("Game Menu:\n1 - Rules\n2 - Create your character\n3 - Fight\n4 - Learn\n5 - Store\n6 - Go to the Principal Menu\nTo exit the game, go to the Main Menu. ;)")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Rules":
			fmt.Print(t3)
			Back()
		case "2", "Create your character":
			p := InitCharacter()
			perso = &p
			fmt.Printf("character creat: %+v\n", p)
			Back()
		case "3", "Fight":
			t5 := Book3()
			if !CheckPerso(perso) {
				break
			}
			Battle(t5, perso)
			Back()
		case "4", "Learn":
			l1, l2 := Learn()
			fmt.Printf("%v\n%v\n", l1, l2)
			Back()
		case "5", "Store":
			if !CheckPerso(perso) {
				break
			}
			Marchand(perso)
			Back()
		case "6", "Go to the Principal Menu":
			t6 := Book4()
			fmt.Println(t6)
			return
		default:
			Invalid()
		}
	}
}

// Displays the inventory and returns the selected index (or -1 if canceled)
func chooseItemIndex(bag []Item) int {
	if len(bag) == 0 {
		fmt.Println("Your inventory is empty.")
		return -1
	}

	fmt.Println("inventory :")
	for i, it := range bag {
		fmt.Printf("%d - %s (%s, power %d)\n", i+1, it.Name, it.Kind, it.Power)
	}
	fmt.Print("Select an item, type ‘0’ to cancel : ")

	var n int
	fmt.Scanln(&n)
	if n <= 0 || n > len(bag) {
		return -1
	}
	return n - 1
}

func applyWeapon(item Item, monster *Monster) {
	min := item.Power / 2
	dmg := rand.Intn(item.Power-min+1) + min
	monster.CurLP -= dmg
	fmt.Printf("used %s and inflict %d damage !\n", item.Name, dmg)
}

func applyHeal(item Item, perso *Character) bool {
	if perso.PvAct >= perso.PvMax-item.Power {
		fmt.Println("Your life is too precious to use this item.")
		return false
	}
	gain := item.Power
	if perso.PvAct+gain > perso.PvMax {
		gain = perso.PvMax - perso.PvAct
	}
	perso.PvAct += gain
	fmt.Printf("used %s and recup %d PV !\n", item.Name, gain)
	return true // consomed
}

func useItem(perso *Character, monster *Monster) {
	idx := chooseItemIndex(perso.Bag)
	if idx < 0 {
		return
	}

	item := perso.Bag[idx]
	switch item.Kind {
	case "weapon":
		applyWeapon(item, monster)
		// remove if weapon used

	case "heal":
		if applyHeal(item, perso) {
			// heal consomed
			perso.Bag = append(perso.Bag[:idx], perso.Bag[idx+1:]...)
		}

	default:
		fmt.Println("inconu object.")
	}
}

func displayStatus(p *Character, m *Monster) {
	fmt.Printf("\n%s : %d/%d PV | %s : %d/%d PV\n",
		p.Name, p.PvAct, p.PvMax,
		m.Name, m.CurLP, m.MaxLP)
}

func askPlayerAction() string {
	fmt.Println("Action :\n1 - Clasic Attack\n2 - used an object in your bag\n3 - Leak")
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func playerAttack(p *Character, m *Monster) {
	base := rand.Intn(21) + 10 // random damage 10..30
	dmg := base + p.Level*2    // option: level-based bonus
	m.CurLP -= dmg
	fmt.Printf("%s attacks and inflicts %d damage !\n", p.Name, dmg)
}

func monsterAttack(p *Character, m *Monster) {
	dmg := rand.Intn(m.AttPt + 1)
	p.PvAct -= dmg
	fmt.Printf("%s attacks and inflicts %d damage !\n", m.Name, dmg)
}

func Battle(t5 Text, perso *Character) {
	fmt.Println("battle start !")
	monster := PickRandomMonster()
	fmt.Printf("%s have %d PV !\n", monster.Name, monster.CurLP)

	for {
		displayStatus(perso, &monster)

		switch askPlayerAction() {
		case "1":
			playerAttack(perso, &monster)
		case "2":
			useItem(perso, &monster)
		case "3":
			fmt.Println("You give up !")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}

		if monster.CurLP <= 0 {
			fmt.Printf("Victory soldat %s !\n", monster.Name)
			return
		}

		monsterAttack(perso, &monster)

		if perso.PvAct <= 0 {
			fmt.Println("Lost battle")
			return
		}
	}
}
