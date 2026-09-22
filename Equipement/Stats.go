package ProjetRED

import personnage "ProjetRED/Personnage"

func RecalculateStats(p *personnage.Character) {
	for _, slot := range []string{p.Weapon, p.Armor, p.Boots, p.Helmet} {
		if item, ok := Items[slot]; ok {
			p.PVMax += item.BonusPV
			p.Strength += item.BonusAtk
			p.Defense += item.BonusDef
			p.Reiki += item.BonusReiki
			p.Spd += item.BonusSpeed
		}
	}

	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}
}
