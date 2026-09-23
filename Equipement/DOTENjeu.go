package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
)

func (p *personnage.Character) UpdateEffects() {
	newEffects := []personnage.StatusEffect{}

	for _, e := range p.Effects {

		// Appliquer l'effet
		p.PV -= e.Damage

		// Clamp PV
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		if p.PV < 0 {
			p.PV = 0
		}

		// Affichage
		if e.Damage > 0 {
			fmt.Println(p.Nom, "subit", e.Damage, "dégâts de", e.Name, "PV :", p.PV)
		} else {
			fmt.Println(p.Nom, "récupère", -e.Damage, "PV grâce à", e.Name, "PV :", p.PV)
		}

		// Réduire le temps restant
		e.TimeLeft -= e.Interval

		// Garder l'effet s’il reste du temps
		if e.TimeLeft > 0 {
			newEffects = append(newEffects, e)
		}
	}

	p.Effects = newEffects
}
