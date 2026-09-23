package ProjetRED

import "fmt"

func DisplayInfo(p Character) string {
	fmt.Println("nom:", p.Nom)
	fmt.Println("class:", p.Classe.Nom)
	fmt.Printf("PV: %d / %d\n", p.PV, p.PVMax)
	fmt.Printf("vous etes lvl: %d\n", p.LVL)
	fmt.Printf("Xp: %.0f\n", p.XP)
	fmt.Println("stat:")
	fmt.Printf("Strength: %d\n", p.Strength)
	fmt.Printf("Defense: %d\n", p.Defense)
	fmt.Printf("Reiki: %d\n", p.Reiki)
	fmt.Printf("Spd: %d\n", p.Spd)
	fmt.Println("vos equipement:")
	fmt.Println("arme:", p.Weapon)
	fmt.Println("casque:", p.Helmet)
	fmt.Println("armure:", p.Armor)
	fmt.Println("bottes:", p.Boots)
	return ""
}
