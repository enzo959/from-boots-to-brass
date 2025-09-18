package main

import (
	"projet-red_from-boots-to-brass/src/auth"
)

func main() {
	t1, t2 := auth.Book1()
	auth.Menu(t1, t2)
}
