package main

import "fmt"

const BLANCC      = "\u001b[0m"
const BLEUCLAIRR  = "\u001b[36m"
const BLEUU       = "\033[34m"
const BLEUFONCEE  = "\033[38;5;18m"
const GRIS        = "\033[90m" 
const OORRANGEE   = "\u001b[38;5;208m"
const JAUNEE      = "\u001b[33m"
const GRISS       = "\033[90m" 
const REDDD       = "\u001b[31m"
const VIOLETT     = "\u001b[35m"
const MARRONN     = "\u001b[38;5;94m"
const OORANGE     = "\u001b[38;5;208m"
const VERTTT      = "\u001b[32m"

func AccessMarchand(b *bell, i *[]inventory) {
	fmt.Println("\n")
	fmt.Println(JAUNEE, "=== Marchand d'exorcisme ===", BLANCC)
	fmt.Println("\n")
	fmt.Println(GRISS,"Âme",BLANCC, "de",BLEUCLAIRR, "spectre",BLANCC, ":", b.Spectre)
	fmt.Println(GRISS,"Âme",BLANCC, "de",BLEUU, "possédé",BLANCC, ":", b.Possede)
	fmt.Println(GRISS,"Âme",BLANCC, "d'",BLEUFONCEE, "hybride",BLANCC, ":", b.Hybride)
	fmt.Println("\n")
	fmt.Println(VERTTT,"🍶 Potions",BLANCC, ":")
	fmt.Println("1) Potion de",REDDD, "vie",BLANCC, "   :",GRISS, "2 âmes",BLANCC, "de", BLEUCLAIRR, "spectres",BLANCC)
	fmt.Println("2) Potion de",BLEUCLAIRR, "mana",BLANCC, "  :",GRISS, "2 âmes",BLANC, "de",BLEUCLAIRR, "spectres",BLANCC)
	fmt.Println("3) Potion de",DARKGREEN, "poison",BLANCC, ":",GRISS, "3 âmes",BLANCC, "de",BLEUCLAIRR, "spectres",BLANCC)
	fmt.Println("\n")
	fmt.Println(VIOLETT,"👻 Résonances : ",BLANCC)
	fmt.Println("4)",BLEUU, "Eau sacrée",BLANCC, "  :",GRISS, "1 âme", BLANCC, "d'",BLEUFONCEE, "hybride",BLANCC)
	fmt.Println("5)",MARRONN, "Roche sacrée",BLANCC, ":", GRISS, "2 âmes",BLANCC, "d'",BLEUFONCEE, "hybride",BLANCC)
	fmt.Println("6)",OORANGE, "Feu sacrée",BLANCC, "  :", GRISS, "3 âmes",BLANCC, "d'",BLEUFONCEE, "hybride",BLANCC)
	fmt.Println("7)",LIGHTGREEN, "Vent sacrée", BLANCC, " :",GRISS, "4 âmes",BLANCC, "d'",BLEUFONCEE, "hybride",BLANCC)
	fmt.Println("\n")
	fmt.Println(MARRONN, "🥼 Matériaux :",BLANCC)
	fmt.Println("8)",BLEUFONCEE, "Spectronyte",BLANCC, ":",GRISS, "1 âme",BLANCC, "de", BLEUCLAIRR, "spectre",BLANCC)
	fmt.Println("0) Retour")
	fmt.Println("\n")
	var saisie int
	var num int
	quit := false

	for {
		if quit { break }
		_, err := fmt.Scanln(&saisie)

		if err != nil {
			fmt.Println("Valeur incorrect.")
			continue
		}
		
		switch saisie {
		case 0 : 
		quit = true
			break
		case 1 :
			if b.Spectre >= 2 {
				num = addInventory("Potion de vie", 1, i)
				b.Spectre -= 2
				fmt.Println("Potion de vie achetée.")
				fmt.Println("Possède :", num)
				fmt.Println("Âme de spectre possédées :", b.Spectre)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme de spectre")
			}
		case 2 :
			if b.Spectre >= 2 {
				num = addInventory("Potion de MANA", 1, i)
				b.Spectre -= 2
				fmt.Println("Potion de mana achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme de spectre possédées :", b.Spectre)			
			} else {
				fmt.Println("Vous n'avez pas assez d'âme de spectre")
			}
		case 3 :
			if b.Spectre >= 3 {
				num = addInventory("Potion de poison", 1, i)
				b.Spectre -= 3
				fmt.Println("Potion de poison achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme de spectre possédées :", b.Spectre)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme de spectre")
			}
		case 4 :
			if b.Hybride >= 1 {
				num = addInventory("Eau sacrée", 1, i)
				b.Hybride -= 1
				fmt.Println("Eau sacrée achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme d'hybride possédées :", b.Hybride)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme d'hybride")
			}
		case 5 : 
			if b.Hybride >= 2 {
				num = addInventory("Roche sacrée", 1, i)
				b.Hybride -= 2
				fmt.Println("Roche sacrée achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme d'hybride possédées :", b.Hybride)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme d'hybride")
			}
		case 6 : 
			if b.Hybride >= 3 {
				num = addInventory("Feu sacrée", 1, i)
				b.Hybride -= 3
				fmt.Println("Feu sacrée achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme d'hybride possédées :", b.Hybride)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme d'hybride")
			}
		case 7 :
			if b.Hybride >= 4 {
				num = addInventory("Vent sacrée", 1, i)
				b.Hybride -= 4
				fmt.Println("Vent sacrée achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme d'hybride possédées :", b.Hybride)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme d'hybride")
			}
		case 8 : 
			if b.Spectre >= 1 {
				 num = addInventory("Spectronyte", 1, i)
				b.Spectre -= 1
				fmt.Println("Spectronyte achetée")
				fmt.Println("Possède :", num)
				fmt.Println("Âme de spectre possédées :", b.Spectre)
			} else {
				fmt.Println("Vous n'avez pas assez d'âme de spectre")
			}
		default :
		}
		if !quit { fmt.Println("Voulez-vous acheter autre chose ?")}
	}
}

func addInventory(s string, n int, i *[]inventory) int {
	for j := range *i {
		if (*i)[j].Nom == s {
			(*i)[j].Nombre += n
			return (*i)[j].Nombre
		}
	}
	return 0
}