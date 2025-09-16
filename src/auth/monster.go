package auth

import "fmt"

type Monster struct {
	name   string
	max_lp int
	cur_lp int
	att_pt int
}

func initGobelin() {
	var m1 Monster // initialisation de ma structure Personnage

	/*
	   Initialisation des attributs de mon personnage p1
	*/
	m1.name = "Training gobelin"
	m1.max_lp = 40
	m1.cur_lp = m1.max_lp
	m1.att_pt = 5

	fmt.Println("Monster LP", m1.name, ":", m1.max_lp)
	fmt.Println("Monster strength", m1.name, ":", m1.att_pt)

	//if m1.mort {
	//  fmt.Println("Vie du personnage", p1.nom, "est mort")
	// } else {
	//     fmt.Println("Vie du personnage", p1.nom, "est vivant")
	// }
}

func initMonster() (Monster, Monster, Monster, Monster, Monster) {

	m2 := Monster{"Xzof", 100, 100, 2}
	m3 := Monster{"Koxy", 80, 100, 3}
	m4 := Monster{"dddd", 80, 100, 2}
	m5 := Monster{"dddd", 70, 100, 2}
	m6 := Monster{"Boss Monster", 100, 100, 10}
	return m2, m3, m4, m5, m6

	//if m.Mort {
	//   fmt.Println("Vie du personnage", p.Nom, "est mort")
	// } else {
	//   fmt.Println("Vie du personnage", p.Nom, "est vivant")
	// }
}
