package ProjetRED

type ItemType string

const (
	Helmet ItemType = "Helmet"
	Armor  ItemType = "Armor"
	Boots  ItemType = "Boots"
	Weapon ItemType = "Weapon"
)

type Item struct {
	Name       string
	BonusPV    int
	BonusAtk   int
	BonusDef   int
	BonusReiki int
	BonusSpeed int
	MaxStack   int
	Type       ItemType
}

type Consumable struct {
	Name     string
	Heal     int
	MaxStack int
}

var Items = map[string]Item{
	"Swordshield":        {Name: "Swordshield", BonusAtk: 5, BonusDef: 10, MaxStack: 1, Type: "Weapon"},
	"Elementalist Rings": {Name: "Elementalist Rings", BonusAtk: 10, BonusReiki: 10, MaxStack: 1, Type: "Weapon"},
	"Bandit Helmet":      {Name: "Bandit Helmet", BonusDef: 10, BonusPV: 15, MaxStack: 1, Type: "Helmet"},
	"Bandit Armor":       {Name: "Bandit Armor", BonusDef: 12, BonusPV: 15, MaxStack: 1, Type: "Armor"},
	"Bandit Boots":       {Name: "Bandit Boots", BonusDef: 7, BonusPV: 10, BonusSpeed: 10, MaxStack: 1, Type: "Boots"},
	"Bandit Spear":       {Name: "Bandit Spear", BonusDef: 2, BonusAtk: 10, MaxStack: 1, Type: "Weapon"},
	"Samourai Helmet":    {Name: "Knight Helmet", BonusDef: 20, BonusPV: 25, MaxStack: 1, Type: "Helmet"},
	"Samourai Armor":     {Name: "Knight Armor", BonusDef: 20, BonusPV: 25, MaxStack: 1, Type: "Armor"},
	"Samourai Boots":     {Name: "Knight Boots", BonusDef: 15, BonusPV: 15, BonusSpeed: 15, MaxStack: 1, Type: "Boots"},
	"Lord Helmet":        {Name: "Lord Helmet", BonusDef: 30, BonusPV: 25, MaxStack: 1, Type: "Helmet"},
	"Lord Armor":         {Name: "Lord Armor", BonusDef: 40, BonusPV: 35, MaxStack: 1, Type: "Armor"},
	"Lord Boots":         {Name: "Lord Boots", BonusDef: 20, BonusPV: 20, BonusSpeed: 20, MaxStack: 1, Type: "Boots"},
	"Lord Battle Axe":    {Name: "Lord Battle Axe", BonusDef: 10, BonusAtk: 20, MaxStack: 1, Type: "Weapon"},
	"Mage Staff":         {Name: "Mage Staff", BonusAtk: 20, MaxStack: 1, Type: "Weapon"},
	"Mage Hood":          {Name: "Mage Hood", BonusReiki: 20, BonusPV: 10, MaxStack: 1, Type: "Helmet"},
	"Mage Robe":          {Name: "Mage Robe", BonusReiki: 30, BonusPV: 10, MaxStack: 1, Type: "Armor"},
	"Mage Boots":         {Name: "Mage Boots", BonusReiki: 10, BonusPV: 10, MaxStack: 1, Type: "Boots"},
	"Katana":             {Name: "Katana", BonusAtk: 10, BonusDef: 10, MaxStack: 1, Type: "Weapon"},
	"Dagger":             {Name: "Dagger", BonusAtk: 15, BonusSpeed: 5, MaxStack: 1, Type: "Weapon"},
}

var HealingPotion = Consumable{
	Name:     "Healing Potion",
	Heal:     50,
	MaxStack: 5,
}
var PoisonDOTPotion = Consumable{
	Name:     "Poison DOT Potion",
	Heal:     0,
	MaxStack: 3,
}
