package auth

import (
	"fmt"
	"time"
)

type Text struct {
	text string
}

func Book2() (Text, Text) {
	t3 := Text{"Rêgles: Vous pouvez faire des combat dans 'Combat'.\nVous pouvez apprendre de nouvelle technique de combat dans 'Learn'.\nVous pouvez aller faire des achat dans 'Marchand'"}
	t4 := Text{"Combat: vous êtes dans l'arène."}
	return t3, t4
}

func Book3() Text {
	t5 := Text{"Menu Attack"}
	return t5
}

func Book4() Text {
	t6 := Text{"Retour au Menu"}
	return t6
}

func Book1() (Text, Text) {
	t1 := Text{"Prologue: \n Dans ce monde, la gloire n’est pas donnée. Elle se prend, balle après balle. \n Je ne suis qu’un soldat parmi des milliers : un matricule, une arme, un ordre. \n On ne choisit pas où l’on commence. On choisit jusqu’où on ira. \n Les tirs résonnent au loin, les hélicoptères déchirent le ciel. Aujourd’hui, ce n’est plus qu’une mission : c’est le premier pas. \n Le champ de bataille forge les hommes… et je compte bien gravir chaque échelon, quel qu’en soit le prix."}
	t2 := Text{"Rêgle: \n Vous pouvez faire des combats dans Game.\nVous pouvez visité le Marchant, qui vous permet de faire des achats."}
	return t1, t2
}

func Clear() { // func clean terminal
	fmt.Print("\033[H\033[2J")
}

func Invalid() {
	fmt.Println("Choisissez une option valide.") // error gestion
	time.Sleep(2 * time.Second)                  // reset after 2 sec
	Clear()                                      // clean terminal
}

func Back() {
	var back string // var back for back to menu
	fmt.Println("\nAppuyez sur 'q' pour revenir en arrière")
	fmt.Scanln(&back)
}

func CheckPerso(perso *Character) bool {
	if perso == nil {
		fmt.Println("Créez d’abord un personnage pour acheter.")
		Back()
		return false
	}
	return true
}
