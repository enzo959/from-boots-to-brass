package game

func Update() (Weapons, Weapons, Weapons, Equipment, Equipment, Equipment, Loot, Loot) {
	w1 := Weapons{"FAMAS", "inflige:", -30, "de pv et coûte", 75, "crédit"}
	w2 := Weapons{"M16", "inflige:", -25, "de pv et coûte", 65, "crédit"}
	w3 := Weapons{"PPSH", "inflige:", -20, "de pv et coûte", 60, "crédit"}
	e1 := Equipment{"Casque", "Redonne :", 50, "pv et coûte :", 20, "crédit"}
	e2 := Equipment{"Pantalon", "Redonne :", 30, "pv et coûte :", 15, "crédit"}
	e3 := Equipment{"Bottes", "Redonne :", 20, "pv et coûte :", 5, "crédit"}
	l1 := Loot{"Bandage", "redonne :", 10, "de pv et coûte :", 5, "crédit"}
	l2 := Loot{"Grenade", "inflige :", -30, "de pv et coûte :", 15, "crédit"}
	return w1, w2, w3, e1, e2, e3, l1, l2
}
