package RED

type Item struct {
	Name       string
	BonusPV    int
	BonusAtk   int
	BonusDef   int
	BonusReiki int
	BonusSpeed int
}
type Consumable struct {
	Name     string
	Heal     int
	MaxStack int
}

var Items = map[string]Item{
	"Swordshield":        {Name: "Swordshield", BonusAtk: 5, BonusDef: 10},
	"Elementalist Rings": {Name: "Elementalist Rings", BonusAtk: 10, BonusReiki: 10},
	"Bandit Helmet":      {Name: "Bandit Helmet", BonusDef: 10, BonusPV: 15},
	"Bandit Armor":       {Name: "Bandit Armor", BonusDef: 12, BonusPV: 15},
	"Bandit Boots":       {Name: "Bandit Boots", BonusDef: 7, BonusPV: 10, BonusSpeed: 10},
	"Bandit Spear":       {Name: "Bandit Spear", BonusDef: 2, BonusAtk: 10},
	"Samourai Helmet":    {Name: "Knight Helmet", BonusDef: 20, BonusPV: 25},
	"Samourai Armor":     {Name: "Knight Armor", BonusDef: 20, BonusPV: 25},
	"Samourai Boots":     {Name: "Knight Boots", BonusDef: 15, BonusPV: 15, BonusSpeed: 15},
	"Lord Helmet":        {Name: "Lord Helmet", BonusDef: 30, BonusPV: 25},
	"Lord Armor":         {Name: "Lord Armor", BonusDef: 40, BonusPV: 35},
	"Lord Boots":         {Name: "Lord Boots", BonusDef: 20, BonusPV: 20, BonusSpeed: 20},
	"Lord Battle Axe":    {Name: "Lord Battle Axe", BonusDef: 10, BonusAtk: 20},
	"Mage Staff":         {Name: "Mage Staff", BonusAtk: 20},
	"Mage Hood":          {Name: "Mage Hood", BonusReiki: 20, BonusPV: 10},
	"Mage Robe":          {Name: "Mage Robe", BonusReiki: 30, BonusPV: 10},
	"Mage Boots":         {Name: "Mage Boots", BonusReiki: 10, BonusPV: 10},
	"Katana":             {Name: "Katana", BonusAtk: 10},
}

var HealingPotion = Consumable{
	Name:     "Healing Potion",
	Heal:     50,
	MaxStack: 5,
}
