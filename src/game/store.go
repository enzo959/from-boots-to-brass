package game

type Loot struct {
	Object    string
	InfoObjet string
	Pv        int
	InfoPv    string
	Price     int
	InfoPrice string
}

type Weapons struct { //create class update
	Name       string
	InfoName   string
	PvAttack   int
	InfoAttack string
	Price      int
	InfoPrice  string
}

type Equipment struct { // create class equipement
	Name      string
	InfoName  string
	PvSup     int
	InfoPvSup string
	Price     int
	InfoPrice string
}

type Skill struct { // create class Skill
	Name          string
	PtAttack      int
	LevelRequiert int
}
