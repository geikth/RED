package projetRED

import (
	ProjetRED "ProjetRED/Personnage"
	"fmt"
	"sort"
	"strings"
)

func DisplayInfo(p ProjetRED.Character) string {
	var sb strings.Builder

	const largeur = 40
	ligne := strings.Repeat("─", largeur)

	fmt.Fprintf(&sb, "╭%s╮\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("%s — %s (Lvl %d)", p.Nom, p.Classe.Nom, p.LVL))
	fmt.Fprintf(&sb, "├%s┤\n", ligne)

	// Barre de PV
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("PV: %d/%d %s", p.PV, p.PVMax, barreDeVie(p.PV, p.PVMax, 15)))
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("XP: %.0f", p.XP))

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Statistiques")
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Force:", p.Strength)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Défense:", p.Defense)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Reiki:", p.Reiki)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Vitesse:", p.Spd)

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Équipement")
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Arme:", p.Weapon)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Casque:", p.Helmet)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Armure:", p.Armor)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Bottes:", p.Boots)

	fmt.Fprintf(&sb, "╰%s╯\n", ligne)

	result := sb.String()
	fmt.Print(result)
	return result
}

// barreDeVie construit une petite barre style [████████░░]
func barreDeVie(pv, pvMax, taille int) string {
	if pvMax <= 0 {
		return ""
	}
	rempli := int(float64(pv) / float64(pvMax) * float64(taille))
	if rempli > taille {
		rempli = taille
	}
	if rempli < 0 {
		rempli = 0
	}
	return "[" + strings.Repeat("█", rempli) + strings.Repeat("░", taille-rempli) + "]"
}

func AccessInventory(p ProjetRED.Character) string {
	var sb strings.Builder

	const largeur = 40
	ligne := strings.Repeat("─", largeur)

	fmt.Fprintf(&sb, "╭%s╮\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Inventaire")

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Objets")
	ecrireSection(&sb, p.Inventory.Items)

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Consommables")
	ecrireSection(&sb, p.Inventory.Consumables)

	fmt.Fprintf(&sb, "╰%s╯\n", ligne)

	result := sb.String()
	fmt.Print(result)
	return result
}

// ecrireSection affiche une map triée par clé, avec un message si elle est vide.
func ecrireSection(sb *strings.Builder, items map[string]int) {
	if len(items) == 0 {
		fmt.Fprintf(sb, "│   %-36s │\n", "Aucun")
		return
	}

	keys := make([]string, 0, len(items))
	for name := range items {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for _, name := range keys {
		ligne := fmt.Sprintf("%s x%d", name, items[name])
		fmt.Fprintf(sb, "│   %-36s │\n", ligne)
	}
}
