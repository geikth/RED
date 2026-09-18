package RED

import (
	"fmt"
)

func (p *character) AddConsumable(item Consumable) {
	current := p.Inventory.Consumables[item.Name]

	if current < item.MaxStack {
		p.Inventory.Consumables[item.Name] = current + 1
		fmt.Println(item.Name, "ajoutée ! Quantité :", current+1)
	} else {
		fmt.Println("Impossible : stack maximum atteint pour", item.Name)
	}
}

func (p *character) UseConsumable(item Consumable) {
	qty := p.Inventory.Consumables[item.Name]

	if qty <= 0 {
		fmt.Println("Tu n'as pas de", item.Name)
		return
	}

	// Empêcher l'utilisation si le joueur est full PV
	if p.PV == p.PVMax {
		fmt.Println("Impossible d'utiliser", item.Name, ": PV déjà au maximum.")
		return
	}

	// Soigne
	p.PV += item.Heal
	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}

	// Détruit une potion
	p.Inventory.Consumables[item.Name] = qty - 1

	fmt.Println(p.Name, "utilise", item.Name, "et récupère", item.Heal, "PV !")
	fmt.Println("PV :", p.PV, "/", p.PVMax)
}
