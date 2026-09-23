package ProjetRED

var Classes = map[string]Classe{
	// sabre / katana
	"Ronin": {
		Nom:      "Ronin",
		PVMax:    30,
		Strength: 20,
		Defense:  10,
		Reiki:    5,
		Spd:      15,
		Weapon:   "Bandit_Spear",
		Helmet:   "Bandit_Helmet",
		Armor:    "Bandit_Armor",
		Boots:    "Bandit_Boots",
	},
	// bouclier combat(donc epée + bouclier)
	"Cuirassé": {
		Nom:      "Cuirassé",
		PVMax:    40,
		Strength: 15,
		Defense:  15,
		Reiki:    5,
		Spd:      5,
		Weapon:   "Swordshield",
		Helmet:   "Samourai_Helmet",
		Armor:    "Samourai_Armor",
		Boots:    "Samourai_Boots",
	},
	// un seul et meme item (bague boucle bijoux)
	"mage spirituel": {
		Nom:      "mage spirituel",
		PVMax:    30,
		Strength: 5,
		Defense:  8,
		Reiki:    30,
		Spd:      15,
		Weapon:   "Elementalist_Rings",
		Helmet:   "Mage_Hood",
		Armor:    "Mage_Robe",
		Boots:    "Mage_Boots",
	},
}
