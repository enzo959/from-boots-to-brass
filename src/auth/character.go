package auth

import (
	"fmt"
	"projet-red_from-boots-to-brass/src/utils"
	"time"
)

type Character struct { // create class character
	name   string
	team   string
	level  int
	pv_max int
	pv_act int
	money  int
	bag    []string
}

type Skill struct { // create class Skill
	name_wp string
	pt_dam  int
}

type Equipment struct { // create class equipement
	boots string
	pants string
	shoes string
}

func displaySkills() Skill {
	p1 := Skill{"test", 10}
	return p1
}

func DisplayInfo() (Character, Character, Character) {
	p1 := Character{"soldat_FR", "France", 0, 100, 100, 0, []string{"FAMAS", "vide", "vide", "vide", "vide"}} // var p1
	p2 := Character{"soldat_US", "U.S.A", 0, 100, 100, 0, []string{"M16", "vide", "vide", "vide", "vide"}}
	p3 := Character{"soldat_RU", "U.R.S.S", 0, 100, 100, 0, []string{"PPSH", "vide", "vide", "vide", "vide"}}
	fmt.Print(displaySkills())
	return p1, p2, p3
}

func CharacterCreation(p1, p2, p3 Character) Character {
	var choiceName string
	var perso Character

	for {
		fmt.Println("Choisissez votre camp :\n1 - France\n2 - U.S.A\n3 - U.R.S.S")
		fmt.Scan(&choiceName)

		switch choiceName {
		case "1", "France": // if user write "1" or "France"
			perso = p1 // perso = p1
		case "2", "U.S.A":
			perso = p2
		case "3", "U.R.S.S":
			perso = p3
		default:
			fmt.Println("Choisissez une option valide.") // error gestion
			time.Sleep(2 * time.Second)                  // reset after 2 sec
			utils.Clear()                                // clean terminal
			continue                                     // restart the code
		}
		break // break the code
	}

	fmt.Println("entre ton nom soldat")
	var newName string
	fmt.Scan(&newName)
	perso.name = newName // attribut new name for perso

	fmt.Printf("\nVous êtes : %s\n", perso.name)
	fmt.Printf("Team : %s | Niveau : %d | PV : %d/%d | Argent : %d\n", //layout
		perso.team, perso.level, perso.pv_act, perso.pv_max, perso.money)
	fmt.Printf("Équipement : %v\n", perso.bag)

	return perso

}
