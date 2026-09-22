package ProjetRED

func ville3() {
	const o = "o"
	const f = "f"

	transportdansvilletrois := map[string]func(){
		j: jack3,
		: ,
	}

	var saisie string

	for {
		fmt.Println("vous êtes dans la ville 3")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack (conseiller avant le desert)")
		fmt.Println("appui sur f pour aller dans la forêt")

		fmt.Scanln(&saisie)

		if saisie == "m" {
			menu()
			return
		}
		if saisie == j || saisie == f {
			break
		}

		fmt.Println("erreur, veuillez entrer une lettre f, j ou m pour le menu")
	}

	transportdansvilletrois[saisie]()
}

}
