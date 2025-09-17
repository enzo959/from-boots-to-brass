package game

func Learn() (Skill, Skill) {
	s1 := Skill{"coup de poing", -5, 2}
	s2 := Skill{"coup de pied", -10, 3}
	return s1, s2
}

func Books() (Text, Text) {
	b1 := Text{"Rêgles: Vous pouvez faire des combat dans 'Combat'.\nVous pouvez apprendre de nouvelle technique de combat dans 'Learn'.\nVous pouvez aller faire des achat dans 'Marchand'"}
	b2 := Text{"Combat: vous êtes dans l'arène."}
	return b1, b2
}
