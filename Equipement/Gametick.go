package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"time"
)

func GameTick(p *personnage.Character) {
	p.UpdateEffects()
	p.UpdateCooldowns()
}

func StartCombat(player *personnage.Character, monster *personnage.Character) {

	for {
		time.Sleep(time.Second)

		GameTick(player)
		GameTick(monster)

		// Si un des deux est mort, on arrête
		if player.PV <= 0 || monster.PV <= 0 {
			break
		}
	}
}
