package game

type Skill struct { // create class Skill
	name_wp string
	pt_dam  int
}

func displaySkills() Skill {
	p1 := Skill{"test", 10}
	return p1
}
