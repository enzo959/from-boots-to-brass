package auth

func Update() (Weapons, Weapons, Weapons, Equipment, Equipment, Equipment, Loot, Loot) {
	w1 := Weapons{"FAMAS", "inflicts:", 30, "of pv and costs", 75, "credit"}
	w2 := Weapons{"M16", "inflicts:", 25, "of pv and costs", 65, "credit"}
	w3 := Weapons{"PPSH", "inflicts:", 20, "of pv and costs", 60, "credit"}
	e1 := Equipment{"Helmet", "Give back:", 50, "of pv and costs", 20, "credit"}
	e2 := Equipment{"Pants", "Give back:", 30, "of pv and costs", 15, "credit"}
	e3 := Equipment{"Boots", "Give back:", 20, "of pv and costs", 5, "credit"}
	l1 := Loot{"Bandage", "Give back:", 10, "of pv and costs", 5, "credit"}
	l2 := Loot{"Grenada", "inflicts:", 30, "of pv and costs", 15, "credit"}
	return w1, w2, w3, e1, e2, e3, l1, l2
}
