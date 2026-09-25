package main

import (
	"fmt"
)

func AccessForgeron(b *bell, i *[]inventory) {
	fmt.Println("\n")
	fmt.Println(ORANGE,"=== Forgeron ===", RESET)
	fmt.Println("\n")
	fmt.Println(GREY,"Âme",RESET, "de",BLEUCLAIR, "spectre",RESET, ":", b.Spectre)
	fmt.Println(GREY,"Âme",RESET, "de",BLEU, "possédé",RESET, ":", b.Possede)
	fmt.Println(GREY,"Âme",RESET, "d'",BLEUFONCE, "hybride",RESET, ":", b.Hybride)
	fmt.Println("\n")
	fmt.Println("⚒️  Armure")
	fmt.Println("1)", BLEUFONCE, "Armure du damné",RESET, ":", GREY, "5 âmes", RESET, "de", BLEU, "possédés", GREY, "10 spectronytes", RESET)
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
				fmt.Println("Vous avez déjà acheté", BLEUFONCE, "l'armure du damné",RESET)
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
				fmt.Println(BLEUFONCE,"Armure du damné",RESET, "achetée.")
				fmt.Println("Possède :", num)
				armure = true
			} else {
				fmt.Println("Vous n'avez pas les ressources nécessaires")
				fmt.Println(GREY,"Âme",RESET, "de possédé : ", b.Possede)
				fmt.Println(BLEUFONCE, "Spectronyte",RESET, ":", index)
			}
			break
		default :
			break
		}
	}
}