package main 

func goblinPattern(tour int,character *MONSTER){
	gobelin := initGoblin()

	dgt := gobelin.Strength
	if tour%3 ==0 {
		dgt= gobelin.Strength * 2
	}
	character.PV* -= dgt

	fmt.Printf("%s inflige à %s %d de dégâts\n", gobelin.NOM, character.NOM, degats)
    fmt.Printf("%d/%d PV\n", character.PV, character.PVMax)
}

func goblinPattern(character *MONSTER){
	Skeleton := initSkeleton()

	dgt := Skeleton.Strength
	
	character.PV* -= dgt

	fmt.Printf("%s inflige à %s %d de dégâts\n", Skeleton.NOM, character.NOM, degats)
    fmt.Printf("%d/%d PV\n", character.PV, character.PVMax)
}

