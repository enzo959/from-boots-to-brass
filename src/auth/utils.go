package auth

import (
	"fmt"
	"time"
)

type Text struct {
	text string
}

func Book2() (Text, Text) {
	t3 := Text{"Rules: You can participate in great battles against great dictators. You can learn new combat tactics. You can increase your artillery."}
	t4 := Text{"You are entering a war zone."}
	return t3, t4
}

func Book3() Text {
	t5 := Text{"Attack Menu"}
	return t5
}

func Book4() Text {
	t6 := Text{"Back to Menu"}
	return t6
}

func Book1() (Text, Text) {
	t1 := Text{"Prologue: \n In this world, glory is not given. It is taken, bullet by bullet. \n I am just one soldier among thousands: a serial number, a weapon, an order. \n We don't choose where we start. We choose how far we will go. Gunfire echoes in the distance, helicopters tear through the sky. Today, it's no longer just a mission: it's the first step. The battlefield forges men... and I intend to climb every rung of the ladder, whatever the cost."}
	t2 := Text{"Rules: \n You can engage in legendary battles. You can upgrade your artillery."}
	return t1, t2
}

func Clear() { // func clean terminal
	fmt.Print("\033[H\033[2J")
}

func Invalid() {
	fmt.Println("Select a valid option.") // error gestion
	time.Sleep(2 * time.Second)           // reset after 2 sec
	Clear()                               // clean terminal
}

func Back() {
	var back string // var back for back to menu
	fmt.Println("\nPress ‘q’ to confirm")
	fmt.Scanln(&back)
}

func CheckPerso(perso *Character) bool {
	if perso == nil {
		fmt.Println("First, create a character to purchase weapons.")
		Back()
		return false
	}
	return true
}
