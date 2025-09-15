package auth

import "fmt"

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

func InitCharacter() {
	fmt.Println("hello world !") //test
	p1 := Character{"soldat_FR", "France", 0, 100, 100, 0, []string{"vide", "vide", "vide", "vide", "vide"}}
    p2 := Character{"soldat_US",   "U.S.A",  0,  100,  100, 0, []string{"vide", "vide", "vide", "vide", "vide"}}
    p3 := Character{"soldat_RU",   "Vert",  0, 100, 100, 0, []string{"vide", "vide", "vide", "vide", "vide"}}
	characters := []Character{p1, p2, p3}
	for _, c := range characters {
    fmt.Printf("%s (team %s) niveau %d\n", c.name, c.team, c.level)
}

func displayinfo() {

}

func characterCreation() {

}

func upgradeInventorySlot() {

}
