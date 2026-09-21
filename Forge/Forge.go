package ProjetRED

import (
	"fmt"
)

type Material struct {
	Name     string
	MaxStack int
}

var Materials = map[string]Material{
	"Fer":     {Name: "Fer", MaxStack: 50},
	"Bois":    {Name: "Bois", MaxStack: 50},
	"Cuir":    {Name: "Cuir", MaxStack: 50},
	"Cristal": {Name: "Cristal", MaxStack: 50},
	"Diamant": {Name: "Diamant", MaxStack: 50},
}

type Recipe struct {
	Result Items
	Cost   map[string]int
}

var ForgeRecipes = map[string]Recipe{
	"La Maxime": {
		Result: Items["La Maxime"],
		Cost: map[string]int{
			"Diamant": 10,
			"Cuir":    5,
			"Fer":     10,
		},
	},
	"Maximilian Boots": {
		Result: Items["Maximilian Boots"],
		Cost: map[string]int{
			"Iron":    10,
			"Leather": 20,
		},
	},
	"Maximilian Armor": {
		Result: Items["Maximilian Armor"],
		Cost: map[string]int{
			"Iron":    20,
			"Crystal": 15,
		},
	},
	"Maximiliann Helmet": {
		Result: Items["Maximilian Helmet"],
		Cost: map[string]int{
			"Diamant": 5,
			"Cuir":    5,
			"Fer":     10,
		},
	},
}

func (p *Character) Forge(itemName string) {
	recipe, ok := ForgeRecipes[itemName]
	if !ok {
		fmt.Println("Recette inconnue :", itemName)
		return
	}

	// Vérifier les matériaux
	for mat, needed := range recipe.Cost {
		if p.Inventory.Materials[mat] < needed {
			fmt.Println("Matériaux insuffisants pour forger", itemName)
			fmt.Println("Il manque :", mat)
			return
		}
	}

	// Retirer les matériaux
	for mat, needed := range recipe.Cost {
		p.Inventory.Materials[mat] -= needed
	}

	// Donner l'objet forgé
	p.GiveItem(recipe.Result)

	fmt.Println(p.Nom, "a forgé :", itemName)
}
