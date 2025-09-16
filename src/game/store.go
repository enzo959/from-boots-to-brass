package game

type armory struct { // create class armory
	name      string
	pt_attack int
}

func displayArmory() armory {
	p1 := armory{"test", 1}
	return p1
}
