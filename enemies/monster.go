package ProjetRED

type LootEntry struct {
	Name string
	Rate int
}

type MONSTER struct {
	NOM      string
	PVMax    int
	PV       int
	PVMAXR   int
	PVR      int
	Strength int
	Defense  int
	Spd      int
	Reiki    int
	Loot     []LootEntry
}

func initGoblin() MONSTER {
	return MONSTER{
		NOM:      "gobelin",
		PVMax:    30,
		PV:       30,
		PVMAXR:   30,
		PVR:      30,
		Strength: 8,
		Defense:  2,
		Spd:      7,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "dague rouillée", Rate: 3},
			{Name: "bourse trouée", Rate: 2},
			{Name: "oreille de gobelin", Rate: 1},
			{Name: "torche éteinte", Rate: 2},
		},
	}
}

func initSkeleton() MONSTER {
	return MONSTER{
		NOM:      "squelette",
		PVMax:    35,
		PV:       35,
		PVMAXR:   35,
		PVR:      35,
		Strength: 10,
		Defense:  4,
		Spd:      5,
		Reiki:    2,
		Loot: []LootEntry{
			{Name: "tibia", Rate: 8},
			{Name: "crâne fissuré", Rate: 5},
			{Name: "épée rouillée", Rate: 20},
		},
	}
}

func initTroll() MONSTER {
	return MONSTER{
		NOM:      "troll",
		PVMax:    80,
		PV:       80,
		PVMAXR:   80,
		PVR:      80,
		Strength: 18,
		Defense:  6,
		Spd:      3,
		Reiki:    0,
		Loot: []LootEntry{
			{Name: "massue en bois", Rate: 6},
			{Name: "peau de troll", Rate: 12},
			{Name: "dent de troll", Rate: 4},
			{Name: "gourdin ébréché", Rate: 5},
		},
	}
}

func initVouivre() MONSTER {
	return MONSTER{
		NOM:      "vouivre",
		PVMax:    55,
		PV:       55,
		PVMAXR:   55,
		PVR:      55,
		Strength: 15,
		Defense:  3,
		Spd:      8,
		Reiki:    12,
		Loot: []LootEntry{
			{Name: "écaille de vouivre", Rate: 15},
			{Name: "griffe de vouivre", Rate: 10},
			{Name: "venin cristallisé", Rate: 20},
		},
	}
}

func initLoupGarou() MONSTER {
	return MONSTER{
		NOM:      "loup-garou",
		PVMax:    65,
		PV:       65,
		PVMAXR:   65,
		PVR:      65,
		Strength: 17,
		Defense:  5,
		Spd:      9,
		Reiki:    4,
		Loot: []LootEntry{
			{Name: "griffe de loup-garou", Rate: 14},
			{Name: "fourrure argentée", Rate: 18},
			{Name: "croc acéré", Rate: 9},
		},
	}
}

func initZombie() MONSTER {
	return MONSTER{
		NOM:      "zombie",
		PVMax:    50,
		PV:       50,
		PVMAXR:   50,
		PVR:      50,
		Strength: 12,
		Defense:  2,
		Spd:      4,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "chair putréfiée", Rate: 5},
			{Name: "os brisé", Rate: 8},
			{Name: "lambeau de tissu", Rate: 5},
			{Name: "anneau rouillé", Rate: 10},
		},
	}
}

func initOrc() MONSTER {
	return MONSTER{
		NOM:      "orc",
		PVMax:    68,
		PV:       68,
		PVMAXR:   68,
		PVR:      68,
		Strength: 19,
		Defense:  7,
		Spd:      4,
		Reiki:    0,
		Loot: []LootEntry{
			{Name: "hache d'orc", Rate: 11},
			{Name: "bouclier cabossé", Rate: 9},
			{Name: "défense d'orc", Rate: 6},
		},
	}
}

func initDragon() MONSTER {
	return MONSTER{
		NOM:      "dragon",
		PVMax:    140,
		PV:       140,
		PVMAXR:   140,
		PVR:      140,
		Strength: 30,
		Defense:  10,
		Spd:      8,
		Reiki:    18,
		Loot: []LootEntry{
			{Name: "écaille de dragon", Rate: 40},
			{Name: "griffe de dragon", Rate: 30},
			{Name: "souffle embouteillé", Rate: 50},
			{Name: "œuf de dragon", Rate: 100},
		},
	}
}

func IsMonsterDead(m *MONSTER) bool {
	return m == nil || m.PV <= 0 || m.PVR <= 0
}
