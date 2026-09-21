package ProjetRED

import "fmt"

type StatusEffect struct {
	Name     string
	Damage   int
	Duration int // en secondes
	Interval int // dégâts toutes les X secondes
	TimeLeft int
}

func (p *character) UseConsumable(item Consumable) {
	qty := p.Inventory.Consumables[item.Name]

	if qty <= 0 {
		fmt.Println("Tu n'as pas de", item.Name)
		return
	}

	// Potion de soin → ne pas utiliser si PV max
	if item.Heal > 0 && p.PV == p.PVMax {
		fmt.Println("Impossible d'utiliser", item.Name, ": PV déjà au maximum.")
		return
	}

	// Potion de poison DOT
	if item.Name == "Poison DOT Potion" {
		effect := StatusEffect{
			Name:     "Poison",
			Damage:   10,
			Duration: 3,
			Interval: 1,
			TimeLeft: 3,
		}
		p.Effects = append(p.Effects, effect)
		fmt.Println(p.Name, "est empoisonné !")
	}

	// Potion de heal instantané
	if item.Heal != 0 {
		p.PV += item.Heal
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		if p.PV < 0 {
			p.PV = 0
		}
	}

	// Détruire une potion
	p.Inventory.Consumables[item.Name] = qty - 1
}
