package main

import "fmt"

const BLANC   = "\u001b[37m"
const ROUGEE  = "\u001b[31m"
const VIOLET  = "\u001b[35m"
const VERTT   = "\u001b[32m"
const ORANGE  = "\u001b[38;5;208m"
const JAUNE   = "\u001b[33m"
const BROWN   = "\u001b[38;5;94m"
const ROSE    = "\033[38;5;201m"

func main() {
	fmt.Println("Bienvenue sur Soul Art Offline !")
	fmt.Println("(entrer LINK pour commencer)")

	var choix string
	fmt.Scan(&choix)

	if choix == "LINK" {
		perso := CreerPersonnage()
		perso.DisplayInfo()
		village(&perso)
	}
}

var b = bell{Spectre : 0, Possede : 0, Hybride : 0}
var ii = []inventory{
	{Nom: "Potion de vie", Type: "soin", Effet: 50, Nombre : 0},
	{Nom: "Potion de MANA", Type: "mana", Effet: 100, Nombre : 0},
	{Nom: "Potion de poison", Type: "degats", Effet: 10, Nombre : 0},
	{Nom: "Eau sacrée", Type: "degats", Effet: 20, Nombre : 0},
	{Nom: "Roche sacrée", Type: "degats", Effet: 40, Nombre : 0},
	{Nom: "Feu sacrée", Type: "degats", Effet: 60, Nombre : 0},
	{Nom: "Vent sacrée", Type: "degats", Effet: 80, Nombre : 0},
	{Nom: "Spectronyte", Type : "", Effet : 0, Nombre : 0},
	{Nom: "Armure du damné", Type : "", Effet : 0, Nombre : 0},
}

func village(p *Personnage) {
	quit := false
	sommeil := 0
	for {
		if quit { break }
		fmt.Println("\n")
		fmt.Println("|||==-Vous êtes dans le village d'exorcistes Saint du Nord-==|||")
		fmt.Println("1.", VERTT, "se battre en forêt", BLANC)
		fmt.Println("2.", JAUNE, "aller chez le Marchand", BLANC)
		fmt.Println("3.", ORANGE, "aller chez le Forgeron", BLANC)
		fmt.Println("4.", BROWN, "acceder à l'inventaire", BLANC)
		fmt.Println("5.", VIOLET, "afficher les statistiques", BLANC)
		fmt.Println("6.", ROSE, "s'entraîner avec le TryGoblin", BLANC)
		fmt.Println("7.", ROUGEE, "aller à l'auberge", BLANC)
		fmt.Println("8.Qui sont-ils ?")
		fmt.Println("9.Quitter")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			foret(p, &b)
		case 2:
			AccessMarchand(&b, &ii)
		case 3:
			AccessForgeron(&b, &ii)
		case 4:
			fmt.Println(MAGENTA, "[--- Inventaire ---]", WHITE)
			afficherinventaire(&ii)
				fmt.Println("\n[--- Boire une potion ---]")
	for j := range ii {
		if ii[j].Nombre > 0 && (ii[j].Type == "soin" || ii[j].Type == "mana") {
			fmt.Println(j+1, ")", ii[j].Nom, "x", ii[j].Nombre)
		}
	}
	fmt.Println("0) Retour")

	var choixPotion int
	fmt.Scan(&choixPotion)

	if choixPotion == 0 {
		break
	}

	idx := choixPotion - 1
	if idx < 0 || idx >= len(ii) || ii[idx].Nombre == 0 || (ii[idx].Type != "soin" && ii[idx].Type != "mana") {
		fmt.Println("Choix invalide.")
		break
	}

	ii[idx].Nombre--

	if ii[idx].Type == "soin" {
		p.stats.HPact += ii[idx].Effet
		if p.stats.HPact > p.stats.HPmax {
			p.stats.HPact = p.stats.HPmax
		}
		fmt.Println("Vous buvez", ii[idx].Nom, "! +", ii[idx].Effet, "HP")
	} else {
		p.stats.MANA += ii[idx].Effet
		fmt.Println("Vous buvez", ii[idx].Nom, "! +", ii[idx].Effet, "MANA")
	}
		case  5:
			fmt.Println(MAGENTA, "[--- Statistiques ---]", WHITE)
			p.DisplayInfo()
		case  6:
			AccesTryGoblin(p, &b)
		case  7:
			sommeil = AccessAuberge(p, sommeil)
		case  8:
			WhoIsBro()
		case  9:
			quit = true
		}
	}
}