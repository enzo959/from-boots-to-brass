package utils

import (
	"fmt"
	"time"
)

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
	fmt.Println("\nAppuyez sur n'importe quelle lettres pour revenir au menu, ou 'q' pour quitter :")
	fmt.Scan(&back)
	if back == "q" {
		return // exit definitiv
	} else {
		Clear()
	}
}
