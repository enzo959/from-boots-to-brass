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

func InitCharacter() {
	var c Character
	p1 := Character{"soldat_FR", "France", 0, 100, 100, 0, []string{"FAMAS", "vide", "vide", "vide", "vide"}}
	p2 := Character{"soldat_US", "U.S.A", 0, 100, 100, 0, []string{"M16", "vide", "vide", "vide", "vide"}}
	p3 := Character{"soldat_RU", "U.R.S.S", 0, 100, 100, 0, []string{"PPSH", "vide", "vide", "vide", "vide"}}
	for {
		fmt.Println("Choisi ton camp:\n1-France\n2-U.S.A\n3-U.R.S.S")
		fmt.Scan(&c.team)
		if c.team == "France" {
			fmt.Print(p1)
			break
		} else if c.team == "U.S.A" {
			fmt.Print(p2)
			break
		} else if c.team == "soldat_RU" {
			fmt.Print(p3)
			break
		} else {
			fmt.Print("Choisissez une team valide.")
			time.Sleep(3 * time.Second)
			utils.Clear()
			continue
		}
	}
}

//func displayinfo() {

//}

//func characterCreation() {

//}

//func upgradeInventorySlot() {

//}
