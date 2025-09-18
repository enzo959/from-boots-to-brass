package auth

import (
	"fmt"
	"math/rand"
)

func Fight(t3, t4 Text) {
	var choiceOption string
	var perso *Character
	for {
		fmt.Println("Menu de jeu:\n1 - Rêgles\n2 - Crée son personage\n3 - Combat\n4 - Learn\n5 - Marchand\n6 - Allé au Menu Principal\nPour quitter le jeu allez au Menu Principal ;)")
		fmt.Scanln(&choiceOption)

		switch choiceOption {
		case "1", "Rêgles":
			fmt.Print(t3)
			Back()
		case "2", "Crée son personage":
			p := InitCharacter()
			perso = &p
			fmt.Printf("Personnage créé: %+v\n", p)
			Back()
		case "3", "Combat":
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
		case "5", "Marchand":
			if !CheckPerso(perso) {
				break
			}
			Marchand(perso)
			Back()
		case "6", "Allé au Menu Principal":
			t6 := Book4()
			fmt.Println(t6)
			return
		default:
			Invalid()
		}
	}
}

// Affiche l'inventaire et renvoie l'index choisi (ou -1 si annuler)
func chooseItemIndex(bag []Item) int {
	if len(bag) == 0 {
		fmt.Println("Inventaire vide.")
		return -1
	}

	fmt.Println("Inventaire :")
	for i, it := range bag {
		fmt.Printf("%d - %s (%s, puissance %d)\n", i+1, it.Name, it.Kind, it.Power)
	}
	fmt.Print("Choisissez un objet (0 pour annuler) : ")

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
	fmt.Printf("Vous utilisez %s et infligez %d dégâts !\n", item.Name, dmg)
}

func applyHeal(item Item, perso *Character) bool {
	if perso.PvAct >= perso.PvMax-item.Power {
		fmt.Println("Vos PV sont trop élevés pour utiliser cet objet.")
		return false
	}
	gain := item.Power
	if perso.PvAct+gain > perso.PvMax {
		gain = perso.PvMax - perso.PvAct
	}
	perso.PvAct += gain
	fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", item.Name, gain)
	return true // consommé
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
		// retirer si arme consommable :
		// perso.Bag = append(perso.Bag[:idx], perso.Bag[idx+1:]...)

	case "heal":
		if applyHeal(item, perso) {
			// soin consommé
			perso.Bag = append(perso.Bag[:idx], perso.Bag[idx+1:]...)
		}

	default:
		fmt.Println("Objet inconnu.")
	}
}

func displayStatus(p *Character, m *Monster) {
	fmt.Printf("\n%s : %d/%d PV | %s : %d/%d PV\n",
		p.Name, p.PvAct, p.PvMax,
		m.Name, m.CurLP, m.MaxLP)
}

func askPlayerAction() string {
	fmt.Println("Actions :\n1 - Attaque basique\n2 - Utiliser objet\n3 - Fuite")
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func playerAttack(p *Character, m *Monster) {
	base := rand.Intn(21) + 10 // dégâts aléatoires 10..30
	dmg := base + p.Level*2    // option : bonus lié au level
	m.CurLP -= dmg
	fmt.Printf("%s attaque et inflige %d dégâts !\n", p.Name, dmg)
}

func monsterAttack(p *Character, m *Monster) {
	dmg := rand.Intn(m.AttPt + 1)
	p.PvAct -= dmg
	fmt.Printf("%s attaque et inflige %d dégâts !\n", m.Name, dmg)
}

func Battle(t5 Text, perso *Character) {
	fmt.Println("⚔️  Un combat commence !")
	monster := PickRandomMonster()
	fmt.Printf("Un %s sauvage apparaît avec %d PV !\n", monster.Name, monster.CurLP)

	for {
		displayStatus(perso, &monster)

		switch askPlayerAction() {
		case "1":
			playerAttack(perso, &monster)
		case "2":
			useItem(perso, &monster)
		case "3":
			fmt.Println("Vous prenez la fuite !")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if monster.CurLP <= 0 {
			fmt.Printf("🎉 Vous avez vaincu %s !\n", monster.Name)
			return
		}

		monsterAttack(perso, &monster)

		if perso.PvAct <= 0 {
			fmt.Println("💀 Vous êtes K.O. !")
			return
		}
	}
}
