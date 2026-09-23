package main

import "fmt"


func AccessAuberge(c *Personnage, sommeil int) int {
	fmt.Println("\n")
	fmt.Println("===🛏️ Auberge 🛏️===")
	fmt.Println("C'est l'heure de rompiche")
	c.stats.HPact = c.stats.HPmax
	c.stats.MANA = c.stats.MANAmax
	fmt.Println("Vous avez récupéré votre santé et mana")
	fmt.Println("\n")
	sommeil++
	switch sommeil {
	case 1 :
		fmt.Println("Vous avez fait votre première nuit")
	case 3 :
		fmt.Println("Vous aimez bien dormir. Mais n'en faites pas une habitude")
	case 6 :
		fmt.Println("Vous dormez beaucoup dernièrement. Vous êtes paresseux, non ?")
	case 9 :
		fmt.Println("Attention ! Si vous dormez à nouveau, vous allez le regretter")
	}
	fmt.Println("\n")
	return sommeil
}