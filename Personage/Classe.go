package character

var Classes = map[string]Classe{
	// sabre / katana
	"Ronin": {
		Nom:      "Ronin",
		PVMax:    51,
		Strength: 27,
		Defense:  18,
		Reiki:    2,
	},
	// bouclier combat(donc epée + bouclier)
	"Cuirasé": {
		Nom:      "Cuirasé",
		PVMax:    58,
		Strength: 16,
		Defense:  22,
		Reiki:    6,
	},
	// un seul et meme item (bague boucle bijoux)
	"mage spirituel": {
		Nom:      "mage spirituel",
		PVMax:    44,
		Strength: 8,
		Defense:  10,
		Reiki:    32,
	},
}
