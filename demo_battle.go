package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func displayCombatMenu(p *personnage.Character, m *enemies.MONSTER) {
	fmt.Println("\n=== MENU DE COMBAT ===")
	fmt.Printf("Joueur : %s  |  PV %d/%d\n", p.Nom, p.PV, p.PVMax)
	fmt.Printf("Ennemi : %s | PV %d/%d\n\n", m.NOM, m.PV, m.PVMax)
	fmt.Println("1. Attaque basique")
	fmt.Println("2. Attaque spéciale")
	fmt.Println("3. Make a Wish")
	fmt.Println("4. Inventaire")
	fmt.Println("5. Défendre")
	fmt.Println("0. Fuir")
}

func makeWishDemo(p *personnage.Character, m *enemies.MONSTER) {
	fmt.Println("\nVous lancez Make a Wish...")
	switch rand.Intn(4) {
	case 0:
		degats := rand.Intn(20) + 5
		p.PV -= degats
		fmt.Printf("Le sort vous revient dessus ! Vous perdez %d PV.\n", degats)
	case 1:
		degats := rand.Intn(15) + 5
		m.PV -= degats
		fmt.Printf("Le sort frappe %s pour %d dégâts !\n", m.NOM, degats)
	case 2:
		soin := 30
		p.PV += soin
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf("Vous récupérez %d PV.\n", soin)
	default:
		fmt.Println("Coup critique ! L'ennemi est éliminé d'un seul coup !")
		m.PV = 0
		m.PVR = 0
	}
}

func main() {
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

	fmt.Println("=== DEMO COMBAT INTERACTIF ===")
	fmt.Printf("Tu affrontes %s !\n", m.NOM)

	reader := bufio.NewReader(os.Stdin)
	for tour := 1; tour <= 12; tour++ {
		fmt.Println("\n--- TOUR", tour, "---")
		displayCombatMenu(p, m)
		fmt.Print("Choix : ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)
		var action int
		if _, err := fmt.Sscanf(choice, "%d", &action); err != nil {
			fmt.Println("Choix invalide.")
			continue
		}

		switch action {
		case 1:
			m.PV -= p.Strength
			fmt.Printf("%s attaque %s pour %d dégâts.\n", p.Nom, m.NOM, p.Strength)
		case 2:
			m.PVR -= p.Reiki
			fmt.Printf("%s utilise une attaque spéciale pour %d dégâts.\n", p.Nom, p.Reiki)
		case 3:
			makeWishDemo(p, m)
		case 4:
			fmt.Println("Inventaire : vide")
		case 5:
			fmt.Println("Vous vous défendez et attendez le prochain coup.")
		case 0:
			fmt.Println("Tu fuis le combat.")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if m.PV <= 0 || m.PVR <= 0 {
			fmt.Printf("%s est vaincu !\n", m.NOM)
			return
		}

		if p.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", p.Nom)
			return
		}

		degats := m.Strength
		p.PV -= degats
		fmt.Printf("%s attaque pour %d dégâts.\n", m.NOM, degats)
		fmt.Printf("%s : PV %d/%d\n", p.Nom, p.PV, p.PVMax)

		if p.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", p.Nom)
			return
		}
	}

	fmt.Println("Le combat se termine sans vainqueur clair.")
}
