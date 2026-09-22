package projetRED

import (
	personnage "ProjetRED/Personnage"
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

// lis la reponse donné depuis le terminal
func ReadChoice(prompt string) (int, bool) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 0, false
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, false
	}

	return value, true
}

// attend que le joeure apuis sur entré
func WaitForReturn() {
	fmt.Println("Appuyez sur Entrée pour continuer...")
	_, _ = fmt.Scanln()
}

// affiche les info du joeur
func DisplayInfo(p personnage.Character) string {
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

// affiche l'inventaire du joueur
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

// fonction renvoie les variables qui seront données au character creator
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
