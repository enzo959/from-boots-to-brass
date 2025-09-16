package auth

type Monster struct {
	name   string
	max_lp int
	cur_lp int
	att_pt int
}

func InitMonster() (Monster, Monster, Monster, Monster, Monster) {
	m1 := Monster{"Training gobelin", 100, 100, 1}
	m2 := Monster{"M2", 100, 100, 1}
	m3 := Monster{"M3", 80, 100, 1}
	m4 := Monster{"M4", 80, 100, 1}
	m5 := Monster{"M5", 70, 100, 1}
	m6 := Monster{"M6", 100, 100, 1}
	return m2, m3, m4, m5, m6
}
