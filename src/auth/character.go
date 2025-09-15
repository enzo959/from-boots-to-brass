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

//type Skill struct { // create class Skill
//	name_wp string
//	pt_dam  int
//}

//type Equipment struct { // create class equipement
//boots string
//pants string
//shoes string
//}

func InitCharacter() (Character, Character, Character) {
	p1 := Character{"soldat_FR", "France", 0, 100, 100, 0, []string{"FAMAS", "vide", "vide", "vide", "vide"}} // var p1
	p2 := Character{"soldat_US", "U.S.A", 0, 100, 100, 0, []string{"M16", "vide", "vide", "vide", "vide"}}
	p3 := Character{"soldat_RU", "U.R.S.S", 0, 100, 100, 0, []string{"PPSH", "vide", "vide", "vide", "vide"}}
	return p1, p2, p3
}

//func displayinfo() {

//}

func CharacterCreation(p1, p2, p3 Character) Character {
	var choice string
	var perso Character

	for {
		fmt.Println("Choisissez votre camp :\n1 - France\n2 - U.S.A\n3 - U.R.S.S")
		fmt.Scan(&choice)

		switch choice {
		case "1", "France":
			perso = p1
		case "2", "U.S.A":
			perso = p2
		case "3", "U.R.S.S":
			perso = p3
		default:
			fmt.Println("Choisissez une option valide.")
			time.Sleep(2 * time.Second)
			utils.Clear()
			continue
		}
		break
	}

	fmt.Println("entre ton nom soldat")
	var newName string
	fmt.Scan(&newName)
	perso.name = newName

	fmt.Printf("\nVous êtes : %s\n", perso.name)
	fmt.Printf("Team : %s | Niveau : %d | PV : %d/%d | Argent : %d\n",
		perso.team, perso.level, perso.pv_act, perso.pv_max, perso.money)
	fmt.Printf("Équipement : %v\n", perso.bag)

	return perso

}

//func upgradeInventorySlot() {

//}
