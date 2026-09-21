package ProjetRED

import (
	"fmt"
	"sort"
)

func AccessInventory(p Character) {
	fmt.Println("Inventaire :")

	fmt.Println("Objets :")
	if len(p.Inventory.Items) == 0 {
		fmt.Println("  Aucun objet")
	} else {
		keys := make([]string, 0, len(p.Inventory.Items))
		for name := range p.Inventory.Items {
			keys = append(keys, name)
		}
		sort.Strings(keys)

		for _, name := range keys {
			fmt.Printf("  - %s : %d\n", name, p.Inventory.Items[name])
		}
	}

	fmt.Println("Consommables :")
	if len(p.Inventory.Consumables) == 0 {
		fmt.Println("  Aucun consommable")
	} else {
		keys := make([]string, 0, len(p.Inventory.Consumables))
		for name := range p.Inventory.Consumables {
			keys = append(keys, name)
		}
		sort.Strings(keys)

		for _, name := range keys {
			fmt.Printf("  - %s : %d\n", name, p.Inventory.Consumables[name])
		}
	}
}
