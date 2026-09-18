package main

func (p *character) RecalculateStats() {
	p.PVMax = 100 + p.Equip.Weapon.BonusPV + p.Equip.Armor.BonusPV
	p.Strength = 10 + p.Equip.Weapon.BonusAtk + p.Equip.Armor.BonusAtk
	p.Defense = 5 + p.Equip.Armor.BonusDef + p.Equip.Weapon.BonusDef

	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}
}
