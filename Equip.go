package main

func (p *character) EquipItem(slot string, item Item) {
	switch slot {
	case "weapon":
		p.Equip.Weapon = item
	case "armor":
		p.Equip.Armor.Chestplate = item
	case "boots":
		p.Equip.Armor.Boots = item
	case "helmet":
		p.Equip.Armor.Helmet = item
	}

	p.RecalculateStats()
}
