package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

func TakePot(p *personnage.Character) {
	effect := personnage.StatusEffect{
		Name:   "Heal",
		Damage: -50,
	}
	p.Effects = append(p.Effects, effect)
	fmt.Println(p.Nom, "se sent revigorer !")
}
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
func ApplyBreadRegen(p *personnage.Character) {
	effect := personnage.StatusEffect{
		Name:   "Regen du Pain",
		Damage: -20,
	}
	p.Effects = append(p.Effects, effect)
	fmt.Println(p.Nom, "mange du pain et commence à se régénérer !")
}
