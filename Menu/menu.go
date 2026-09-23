package projetRED

import (
	Forge "ProjetRED/Forge"
	personnage "ProjetRED/Personnage"
	MarchandPkg "ProjetRED/marchand"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func StartMenu() {
	fmt.Println("\n=== MENU CREATION ===")
	fmt.Println("1. Crée un nouveau Personage")
	fmt.Println("0. Quitter")
	choice, reponse := ReadChoice("Votre choix : ")
	if !reponse {
		fmt.Println("Choix invalide !")
		StartMenu()
		return
	}
	switch choice {
	case 1:
		nom, classe := CreerPerso()
		Player := personnage.CharacterCreation(nom, classe)
		MainMenu(Player)
	case 0:
		fmt.Println("Au Revoir !")
		return
	default:
		fmt.Println("Choix invalide !")
	}
}

func Marchand(p personnage.Character) {
	fmt.Println("\n=== MARCHAND ===")
	MarchandPkg.Marchand()
	_ = p
}

func Forgeron(p personnage.Character) {
	fmt.Println("\n=== FORGERON ===")
	fmt.Println("Aucune recette sélectionnée pour le moment.")
	_ = p
	_ = Forge.Forge
}

func MainMenu(p personnage.Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("0. Quitter")

		choice, reponse := ReadChoice("Votre choix : ")
		if !reponse {
			fmt.Println("Choix invalide !")
			continue
		}

		switch choice {
		case 1:
			DisplayInfo(p)
			WaitForReturn()
		case 2:
			AccessInventory(p)
			WaitForReturn()
		case 3:
			Marchand(p)
		case 4:
			Forgeron(p)
		case 0:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

func ReadChoice(prompt string) (int, bool) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return 0, false
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return 0, false
	}
	value, err := strconv.Atoi(line)
	if err != nil {
		return 0, false
	}
	return value, true
}

func WaitForReturn() {
	fmt.Println("Appuyez sur Entrée pour continuer...")
	_, _ = fmt.Scanln()
}

func DisplayInfo(p personnage.Character) string {
	var sb strings.Builder
	const largeur = 40
	ligne := strings.Repeat("─", largeur)
	fmt.Fprintf(&sb, "╭%s╮\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("%s — %s (Lvl %d)", p.Nom, p.Classe.Nom, p.LVL))
	fmt.Fprintf(&sb, "├%s┤\n", ligne)
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

func AccessInventory(p personnage.Character) string {
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

func CreerPerso() (string, personnage.Classe) {
	fmt.Print("Quel est votre nom ? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nom := strings.TrimSpace(scanner.Text())
	fmt.Println("Choisissez le numero d'une classe :")
	fmt.Println("1. Ronin")
	fmt.Println("2. Cuirassé")
	fmt.Println("3. Mage spirituel")
	choice, ok := ReadChoice("Votre choix : ")
	if !ok {
		fmt.Println("Choix invalide, classe par défaut : Ronin")
		return nom, personnage.Classes["Ronin"]
	}
	switch choice {
	case 1:
		return nom, personnage.Classes["Ronin"]
	case 2:
		return nom, personnage.Classes["Cuirasé"]
	case 3:
		return nom, personnage.Classes["mage spirituel"]
	default:
		fmt.Println("Choix invalide, classe par défaut : Ronin")
		return nom, personnage.Classes["Ronin"]
	}
}
