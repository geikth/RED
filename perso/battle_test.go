package ProjetRED

import (
	"strings"
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestRunDemoBattle(t *testing.T) {
	p := &personnage.Character{
		Nom:      "Alex",
		PVMax:    200,
		PV:       200,
		Strength: 40,
		Reiki:    25,
		Spd:      50,
	}
	m := &enemies.MONSTER{
		NOM:      "Gobelin",
		PVMax:    150,
		PV:       150,
		PVMAXR:   150,
		PVR:      150,
		Strength: 18,
		Spd:      10,
	}

	result := RunDemoBattle(p, m, []int{1, 2, 1, 1, 1, 1})
	if result == "" {
		t.Fatal("le combat de démonstration ne doit pas être vide")
	}
	if !strings.Contains(result, "Alex") && !strings.Contains(result, "Gobelin") {
		t.Fatalf("résultat du combat invalide : %q", result)
	}
	if p.PV <= 0 {
		t.Fatal("le joueur ne doit pas finir mort dans ce test de combat")
	}
}
