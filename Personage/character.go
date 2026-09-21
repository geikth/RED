package ProjetRED

type character struct {
	Nom       string
	Classe    Classe
	LVL       int
	XP        float64
	PVMax     int
	PV        int
	Strength  int
	Defense   int
	Reiki     int
	Spd       int
	Weapon    string
	Helmet    string
	Armor     string
	Boots     string
	Inventory Inventory
	Effects   []StatusEffect
}

type Classe struct {
	Nom      string
	PVMax    int
	Strength int
	Defense  int
	Reiki    int
	Spd      int
}

func CharacterCreation(nom string, classe Classe) character {
	return character{
		Nom:      Capitalize(nom),
		Classe:   classe,
		LVL:      1,
		XP:       0,
		PVMax:    classe.PVMax,
		PV:       classe.PVMax, // PV de départ = PVMax
		Strength: classe.Strength,
		Defense:  classe.Defense,
		Reiki:    classe.Reiki,
		Spd:      classe.Spd,
	}
}

type Inventory struct {
	Items       map[string]int
	Consumables map[string]int
}
