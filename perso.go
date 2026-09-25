package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Personnage struct {
	nom    string
	classe string
	invent []string
	stats  Stats
}

type Race struct {
	Nom string
}

type Stats struct {
	HPmax       int
	HPact       int
	DGT         int
	MANA        int
	MANAmax     int
	itemSpecial string
}

func couleurRace(race string) string {
	switch race {
	case "Elfe":
		return BROWN
	case "Vampire":
		return RED
	case "Humain":
		return BEIGE
	case "Slime":
		return GREEN
	default:
		return ""
	}
}

func (c *Personnage) Init(nom string, classe string, invent []string, stats Stats) {
	c.nom = nom
	c.classe = classe
	c.invent = invent
	c.stats = stats
	c.invent = append(c.invent, stats.itemSpecial)
}

func (c Personnage) DisplayInfo() {
	fmt.Println("==================================================================================================================")
	fmt.Println("Nom du perso:", c.nom)
	fmt.Println("Classe du perso:", couleurRace(c.classe)+c.classe+RESET)
	fmt.Println("-------------------")
	fmt.Println("Stats du perso:")
	fmt.Println(GREEN,    "  HpMax",RESET, "du perso:   ",GREEN, c.stats.HPmax,RESET)
	fmt.Println(GREEN,    "  HpActuel",RESET, "du perso:",GREEN, c.stats.HPact,RESET)
	fmt.Println(RED,      "  DGT",RESET, ":             ",RED, c.stats.DGT,RESET)
	fmt.Println(BLEUCLAIR,"  MANAmax",RESET, ":         ",BLEUCLAIR, c.stats.MANAmax,RESET)
	fmt.Println(BLEUCLAIR,"  MANA",RESET, ":            ",BLEUCLAIR, c.stats.MANA,RESET)
	fmt.Println("-------------------")
	fmt.Println(BROWN,"Inventaire",RESET, "du perso:", c.invent)
}

func getStats(race string) Stats {
	switch race {
	case "Elfe":
		return Stats{HPmax: 6, HPact: 3, DGT: 6, MANAmax: 6, MANA: 6, itemSpecial: "Brindille 18$"}
	case "Vampire":
		return Stats{HPmax: 80, HPact: 40, DGT: 12, MANAmax: 100, MANA: 100, itemSpecial: "Griffusion"}
	case "Humain":
		return Stats{HPmax: 100, HPact: 50, DGT: 10, MANAmax: 100, MANA: 100, itemSpecial: "Purificateur"}
	case "Slime":
		return Stats{HPmax: 120, HPact: 60, DGT: 8, MANAmax: 100, MANA: 100, itemSpecial: "Gloubs"}
	default:
		return Stats{HPmax: 100, HPact: 50, DGT: 10, MANAmax: 100, MANA: 100, itemSpecial: ""}
	}
}

func CreerPersonnage() Personnage {
	var c Personnage
	var nom string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(" ")
		fmt.Println("Choisissez le nom de votre personnage : ")
		nom, _ = reader.ReadString('\n')
		nom = strings.TrimSpace(nom)
		if len(nom) == 0 {
			fmt.Println("Ce nom est invalide")
			continue
		}
		valide := true
		for _, caractere := range nom {
			if !unicode.IsLetter(caractere) {
				valide = false
				break
			}
		}
		if !valide {
			fmt.Println("Ce nom est invalide")
			continue
		}
		nom = strings.ToLower(nom)
		runes := []rune(nom)
		runes[0] = unicode.ToUpper(runes[0])
		nom = string(runes)
		break
	}

	elfe := Race{"Elfe"}
	vampire := Race{"Vampire"}
	humain := Race{"Humain"}
	slime := Race{"Slime"}

	races := []Race{elfe, vampire, humain, slime}

	fmt.Println("Choisissez votre race :")
	for i, race := range races {
		fmt.Println(i+1, ". "+couleurRace(race.Nom)+race.Nom+RESET) 
	}

	var numero int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&numero)

	raceChoisie := races[numero-1]
	stats := getStats(raceChoisie.Nom)
	fmt.Println("Vous avez choisi :", couleurRace(raceChoisie.Nom)+raceChoisie.Nom+RESET)
	c.Init(nom, raceChoisie.Nom, []string{}, stats)
	return c
}