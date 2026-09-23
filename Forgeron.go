package main

import (
	"fmt"
)

func AccessForgeron(b *bell, i *[]inventory) {
	fmt.Println("\n")
	fmt.Println("=== Forgeron ===")
	fmt.Println("Âme de spectre : ", b.Spectre)
	fmt.Println("Âme de possédé : ", b.Possede)
	fmt.Println("Âme d'hybride : ", b.Hybride)
	fmt.Println("⚒️  Armure")
	fmt.Println("1) Armure du damné : 5 âmes de possédés, 10 spectronytes")
	fmt.Println("0) Retour")
	fmt.Println("\n")
	var saisie int
	var num int
	quit := false
	armure := false
	for {
		if quit { break }
		_, err := fmt.Scanln(&saisie)

		if err != nil {
			fmt.Println("Valeur incorrect.")
			continue
		}


		index := 0
		switch saisie {
		case 0 : 
			quit = true
			break
		case 1 :
			if armure { 
				fmt.Println("Vous avez déjà acheté l'armure du damné")
				break
			}
			for j := range *i {
				if (*i)[j].Nom == "Spectronyte" {
					index = (*i)[j].Nombre
				}
			}
			if b.Possede >= 5 && index >= 10 {
				num = addInventory("Armure du damné", 1, i)
				b.Possede -= 5
				for j := range *i {
					if (*i)[j].Nom == "Spectronyte" {
						(*i)[j].Nombre -= 10
					}
				}
				fmt.Println("Armure du damné achetée.")
				fmt.Println("Possède :", num)
				armure = true
			} else {
				fmt.Println("Vous n'avez pas les ressources nécessaires")
				fmt.Println("Âme de possédé : ", b.Possede)
				fmt.Println("Spectronyte : ", index)
			}
			break
		default :
			break
		}
	}
}