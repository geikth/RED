package ProjetRED

import (
	"fmt"
	"strconv"
)

func jack1(){
	var saisie string
	fmt.Println("bonjours apprenti voyageur!")
	wait(0.5)
	fmt.Print("je suis jack ton fidèle amis donneur de quête...")
	wait(0.5)
	fmt.Println("chiante il faut se l'avouer")
	fmt.Println("mais que serai se monde sans ?")
	fmt.Println("c'est alors pour cela que j'ai besoin de ton aide!")
	fmt.Println("je voudrais que tu tue le troll des forêts?")
	fmt.Println("en échange d'une récompense bien sur")
	fmt.Println("si tu n'accepte pas il y aura des concéquances désastreuse sur cette ville!")
	fmt.Pintln("alors si tu accepte ma quête appui sur o, sinon n")
	fmt.Scanln(&saisie)
	if saisie == "m" {
			menu()
			return
		}
	choixjoueurj1, err := strconv.Atoi(saisie)
	if err != nil || choixjoueurj1 != o || choixjoueurj1 != n {
		fmt.Println("erreur, veuillez entrer une lettre o, n ou m pour le menu")
		continue
	}
	if 
}