package ProjetRED

import (
	"fmt"
	"strconv"
)

func ville1() {
	var saisie string
	transportdansvilleun := []func(){foret, guilde}
	fmt.Scanln(&saisie)
	fmt.Println("vous êtes dans la ville 1")
	fmt.Println("appui sur m pour le menu")
	fmt.Println("appui sur g pour aller dans la guilde (conseiller avant la forêt)")
	fmt.Println("appui sur f pour aller dans la forêt")
	
	if saisie == "m" {
			menu()
			return
		}
	choixdansvilleun, err := strconv.Atoi(saisie)
	if err != nil || choixdansvilleun != g || choixdansvilleun != f {
		fmt.Println("erreur, veuillez entrer une lettre f, g ou m pour le menu")
		continue
	}
	transportdansvilleun[choixdansvilleun-1]()
	return
	