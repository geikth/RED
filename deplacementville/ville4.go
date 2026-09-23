package ProjetRED

import (
	"fmt"
)

func Ville4() {
	const j = "j"
	const p = "p"

	transportdansvilletrois := map[string]func(){
		j: Jack4,
		p: Plaine,
	}
	var saisie string

	for {
		fmt.Println("vous êtes dans la ville 3")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack (conseiller avant la plaine)")
		fmt.Println("appui sur p pour aller dans la plaine")

		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}
		if saisie == j || saisie == p {
			break
		}

		fmt.Println("erreur, veuillez entrer une lettre p, j ou m pour le menu")
	}

	transportdansvilletrois[saisie]()
}
