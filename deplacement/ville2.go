package ProjetRED

import (
	"fmt"
	"strconv"
)
func ville2() {
		var saisie string
	transportdansvilleun := []func(){grotte, jack2, marchand}
	fmt.Scanln(&saisie)
	fmt.Println("vous êtes dans la ville 2")
	fmt.Println("vous avez débloquer l'interaction avec le marchand")
	fmt.Println("appuie sur a pour acceder au marchand")
	fmt.Println("appui sur m pour le menu")
	fmt.Println("appui sur j pour aller parler a jack (conseiller avant la grotte)")
	fmt.Println("appui sur g pour aller dans la grotte")
	
	if saisie == "m" {
			menu()
			return
		}
	choixdansvilleun, err := strconv.Atoi(saisie)
	if err != nil || choixdansvilledeux != j || choixdansvilledeux != f || choixdansvilledeux != a {
		fmt.Println("erreur, veuillez entrer une lettre g, j, a ou m ")
		continue
	}
	transportdansvilledeux[choixdansvilledeux-1]()
	return
}
