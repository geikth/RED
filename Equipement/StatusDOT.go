package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

func ApplyPoisonDOT(p *personnage.Character) {
	effect := personnage.StatusEffect{
		Name:     "Poison",
		Damage:   10,
		Duration: 3,
		Interval: 1,
		TimeLeft: 3,
	}
	p.Effects = append(p.Effects, effect)
	fmt.Println(p.Nom, "est empoisonné !")
}
