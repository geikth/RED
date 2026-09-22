package main

import (
	Equipement "ProjetRED/Equipement"
	Personnage "ProjetRED/Personnage"
)

func main() {
	jean := Personnage.CharacterCreation("jean", Personnage.Classes["Ronin"])
	Personnage.DisplayInfo(jean)
	Equipement.AddItem(jean, Equipement.Items["Swordshield"])
	Personnage.AccessInventory(jean)
}
