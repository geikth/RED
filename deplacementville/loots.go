package ProjetRED

import (
	"fmt"
	"math/rand"
)

type Item struct {
	Nom  string
	Prix int
}

type Monstre struct {
	Nom  string
	Loot []Item
}

func monstresloot() {
	monstres := []Monstre{
		{
			Nom: "troll",
			Loot: []Item{
				{"massue en bois", 6},
				{"peau de troll", 12},
				{"dent de troll", 4},
				{"gourdin ébréché", 5},
			},
		},
		{
			Nom: "vouivre",
			Loot: []Item{
				{"écaille de vouivre", 15},
				{"griffe de vouivre", 10},
				{"venin cristallisé", 20},
			},
		},
		{
			Nom: "gobelin",
			Loot: []Item{
				{"dague rouillée", 3},
				{"bourse trouée", 2},
				{"oreille de gobelin", 1},
				{"torche éteinte", 2},
			},
		},
		{
			Nom: "loup-garou",
			Loot: []Item{
				{"griffe de loup-garou", 14},
				{"fourrure argentée", 18},
				{"croc acéré", 9},
			},
		},
		{
			Nom: "zombie",
			Loot: []Item{
				{"chair putréfiée", 5},
				{"os brisé", 8},
				{"lambeau de tissu", 5},
				{"anneau rouillé", 10},
			},
		},
		{
			Nom: "orc",
			Loot: []Item{
				{"hache d'orc", 11},
				{"bouclier cabossé", 9},
				{"défense d'orc", 6},
			},
		},
		{
			Nom: "squelette",
			Loot: []Item{
				{"tibia", 8},
				{"crâne fissuré", 5},
				{"épée rouillée", 20},
			},
		},
		{
			Nom: "dragon",
			Loot: []Item{
				{"écaille de dragon", 40},
				{"griffe de dragon", 30},
				{"souffle embouteillé", 50},
				{"œuf de dragon", 100},
			},
		},
		{
			Nom: "Ver des Sables",
			Loot: []Item{
				{"Ecaille de Ver", 25},
				{"Dent de Ver", 15},
				{"Sable Cristallisé", 40},
				{"Venin de Sable", 30},
			},
		},
		{
			Nom: "Onryō",
			Loot: []Item{
				{"Fragment d'Âme Vengeresse", 90},
				{"Amulette de Malédiction", 70},
				{"Lambeau de Kimono Blanc", 35},
				{"Larme de Rancune", 45},
			},
		},
	}

	for _, m := range monstres {
		item := m.Loot[rand.Intn(len(m.Loot))]
		fmt.Printf("%s a lâché : %s (%d pièces)\n", m.Nom, item.Nom, item.Prix)
	}
}
