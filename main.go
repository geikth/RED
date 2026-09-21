package main

import (
	Equipement "ProjetRED/Equipement"
	Menu "ProjetRED/Menu"
	Personnage "ProjetRED/Personnage"
)

func main() {
	jean := Personnage.CharacterCreation("jean", Personnage.Classes["mage spirituel"])
	Equipement.AddItem(jean, Equipement.Items["Swordshield"])
	Menu.DisplayInfo(jean)
	Menu.AccessInventory(jean)
}
