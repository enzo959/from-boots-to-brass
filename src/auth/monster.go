package auth

import "math/rand"

type Monster struct {
	Name  string
	MaxLP int
	CurLP int
	AttPt int // attaque de base
}

func DisplayMonster() []Monster {
	return []Monster{
		{"soldat_entrainement", 100, 100, 10},
		{"Hitler", 200, 200, 25},
		{"Mussolini", 120, 120, 15},
		{"Franco", 150, 150, 20},
		{"Pétain", 180, 180, 18},
	}
}

func PickRandomMonster() Monster {
	monsters := DisplayMonster()
	// pondérations (ex : plus petit = plus rare)
	weights := []int{40, 5, 20, 15, 20}
	total := 0
	for _, w := range weights {
		total += w
	}
	r := rand.Intn(total)
	for i, w := range weights {
		if r < w {
			return monsters[i]
		}
		r -= w
	}
	return monsters[0] // fallback
}
