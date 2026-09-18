package RED

func (p *character) RecalculateStats() {
	p.PVMax += p.Equip.Weapon.BonusPV + p.Equip.Armor.BonusPV
	p.Strength += p.Equip.Weapon.BonusAtk + p.Equip.Armor.BonusAtk
	p.Defense += p.Equip.Armor.BonusDef + p.Equip.Weapon.BonusDef
	p.Reiki += p.Equip.Armor.BonusReiki + p.Equip.Weapon.BonusReiki
	p.Spd += p.Equip.Armor.BonusSpeed + p.Equip.Weapon.BonusSpeed

	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}
}
