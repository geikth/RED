package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
)

func TestEquipItemAddsBonusesFromAllEquippedSlots(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classe{
		Nom:      "Test",
		PVMax:    100,
		Strength: 10,
		Defense:  5,
		Reiki:    2,
		Spd:      3,
	})
	weapon := Item{Name: "Test Weapon", BonusAtk: 7, BonusDef: 4, Type: Weapon}
	armor := Item{Name: "Test Armor", BonusPV: 25, BonusReiki: 6, Type: Armor}
	Items[weapon.Name] = weapon
	Items[armor.Name] = armor
	defer delete(Items, weapon.Name)
	defer delete(Items, armor.Name)

	EquipItem(&player, "weapon", weapon)
	EquipItem(&player, "armor", armor)

	if player.Strength != 17 || player.Defense != 9 || player.Reiki != 8 {
		t.Fatalf("unexpected stats: strength=%d defense=%d reiki=%d", player.Strength, player.Defense, player.Reiki)
	}
	if player.PVMax != 125 || player.PV != 125 {
		t.Fatalf("expected PV 125/125, got %d/%d", player.PV, player.PVMax)
	}
}
