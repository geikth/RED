package ProjetRED

import (
	"fmt"
	"strconv"
)

func ville1() {
	var saisie string
	transportdansvilleun := []func(){foret, jack1}
	fmt.Scanln(&saisie)
	fmt.Println("vous êtes dans la ville 1")
	fmt.Println("appui sur m pour le menu")
	fmt.Println("appui sur j pour aller parler a jack (conseiller avant la forêt)")
	fmt.Println("appui sur f pour aller dans la forêt")
	
	if saisie == "m" {
			menu()
			return
		}
	choixdansvilleun, err := strconv.Atoi(saisie)
	if err != nil || choixdansvilleun != j || choixdansvilleun != f {
		fmt.Println("erreur, veuillez entrer une lettre f,  ou m pour le menu")
		continue
	}
	transportdansvilleun[choixdansvilleun-1]()
	return
}
