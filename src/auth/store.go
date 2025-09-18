package auth

type Loot struct {
	Name      string
	InfoObjet string
	PvAttack  int
	InfoPv    string
	Price     int
	InfoPrice string
}

type Weapons struct { //create struct update
	Name       string
	InfoName   string
	PvAttack   int
	InfoAttack string
	Price      int
	InfoPrice  string
}

type Equipment struct { // create struct equipement
	Name      string
	InfoName  string
	PvSup     int
	InfoPvSup string
	Price     int
	InfoPrice string
}

type Skill struct { // create struct Skill
	Name          string
	PtAttack      int
	LevelRequiert int
}

type StoreItem struct { // create struct StoreItem
	Name  string
	Price int
}
