package main

import (
	"os"
	"fmt"
	"math/rand"
)

const RESET      = "\u001b[0m"
const RED        = "\u001b[31m"
const YELLOW     = "\u001b[33m"
const BLEUCLAIR  = "\u001b[36m"
const BLEU       = "\033[34m"
const BLEUFONCE  = "\033[38;5;18m"
const LIGHTGREEN = "\033[38;5;120m"
const GREEN      = "\u001b[32m"
const DARKGREEN  = "\033[38;5;22m"
const PURPLE     = "\u001b[35m"
const BROWN      = "\u001b[38;5;94m"
const ORANGE     = "\u001b[38;5;208m"
const GREY       = "\033[90m" 
const ROSE       = "\033[38;5;201m"
const BEIGE   = "\u001b[38;5;180m"

type Monstre struct {
	Nom string
	HP  int
	Dgt int
	Boss bool
}

func foret(p *Personnage, b *bell) {
	fmt.Println("\n")
	fmt.Println("[-Vous etes dans la forêt,faites attention rebroussez chemin avant de mourir!-]")
	fmt.Println("1.", RED, "Se battre", RESET)
	fmt.Println("2.", YELLOW, "Rebrousser chemin", RESET)

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		fmt.Println("\n")
		fmt.Println("[-Qui voulez-vous affronter?-]")
		fmt.Println("1.", BLEUCLAIR, "Spectre", RESET)
		fmt.Println("2.", BLEU, "Possédé", RESET)
		fmt.Println("3.", BLEUFONCE, "Hybride", RESET)

		var choixMonstre int
		fmt.Scan(&choixMonstre)

		var m Monstre
		switch choixMonstre {
		case 1:
			m = Monstre{Nom: "Spectre", HP: 50, Dgt: 5}
		case 2:
			m = Monstre{Nom: "Possédé", HP: 75, Dgt: 10}
		case 3:
			m = Monstre{Nom: "Hybride", HP: 100, Dgt: 15}
		}
		combat(p, m, b)
	}
}

func combat(p *Personnage, m Monstre, b *bell) {
	fmt.Println("\n")
	fmt.Println(YELLOW, "[Le DU-DU-DU-DUEL contre un", m.Nom, "commence!]", RESET)

	tour := 0 
	
	for p.stats.HPact > 0 && m.HP > 0 {
		tour++ 
		fmt.Println("\n ⚔️  |=== Tour de combat", tour," ===| ⚔️")
		fmt.Println("Vous:",GREEN, p.stats.HPact, "HP",RESET, " |", m.Nom, ":", GREEN, m.HP, "HP", RESET)
		fmt.Println(BLEUCLAIR, p.stats.MANA, "    MANA", RESET, "restant")
		fmt.Println("1.", RED, "Attaquer", RESET)
		fmt.Println("2.", PURPLE, "Résonances", RESET)
		fmt.Println("3.", BROWN, "Inventaire", RESET)
		fmt.Println("4.  Fuir")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			m.HP -= p.stats.DGT
			fmt.Println("Vous infligez", RED, p.stats.DGT, "degats", RESET, "à", m.Nom,"!")
			if m.HP <= 0 {
				fmt.Println(m.Nom,"est vaincu !")
				finCombat(p, m, b, tour)
				return
			}
		case 2:
			fmt.Println("1.", BROWN, "Coup de Poing", RESET, "(",RED,"8 dégâts",RESET,",",BLEUCLAIR,"8 MANA",RESET,")")
			fmt.Println("2.", ORANGE, "Boule de Feu", RESET, "(",RED,"18 dégâts",RESET,",",BLEUCLAIR,"18 MANA",RESET,")")
			fmt.Println("3.", RED, "Dévotion Vampirique", RESET, "(vide tout ton",BLEUCLAIR,"MANA",RESET, "en", RED, "dégâts", RESET, ",", BLEUCLAIR, "150 MANA", RESET, "min)")
			fmt.Println("4.", GREEN, "ROMAIN COTTAR", RESET, "(",GREEN, "1 dégats",GREEN, ",",GREEN,"1 MANA",RESET,")")
			fmt.Println("5.  Retour")

			var choixComp int
			fmt.Scan(&choixComp)

			switch choixComp {
			case 1:
				if p.stats.MANA >=  8 {
					p.stats.MANA -=  8
					m.HP -=  8
					fmt.Println("Vous lancez", BROWN, "Coup de Poing", RESET, "!", RED, "8 dégâts", RESET, "à", m.Nom,"!")
				} else {
					fmt.Println("Pas assez de", BLEUCLAIR, "MANA", RESET,  " !")
				}
			case 2:
				if p.stats.MANA >=  18 {
					p.stats.MANA -=  18
					m.HP -= 18
					fmt.Println("Vous lancez", ORANGE, "Boule de Feu", RESET, "!", RED, "18 dégâts", RESET, "à", m.Nom,"!")
				} else {
					fmt.Println("Pas assez de", BLEUCLAIR, "MANA", RESET,  " !")
				}
			case 3:
				if p.stats.MANA >=  150 {
					m.HP -= p.stats.MANA
					fmt.Println("Vous libérez toute votre", YELLOW, "énergie", RESET, "!", RED, p.stats.MANA,"dégâts", RESET, "à", m.Nom,"!")
					p.stats.MANA = 0
				} else {
					fmt.Println("Il vous faut au moins", BLEUCLAIR, "150 MANA", RESET, "pour cette résonance !")
				}
			case 4:
				if p.stats.MANA >= 1 {
				 m.HP -= 1
				fmt.Println(RED,"Vous envoyez", YELLOW,"ROMAIN COTTAR",RED, "pour qu'il lêche le CHIBRAX de l'ennemie !", RESET)
			}else {
				fmt.Println(RED,"!!!-Il faut au moins 1 de MANA pour envoyer ce gros suceur de", YELLOW, "ROMAIN COTTAR-!!!", RESET)
			}
			case 5:
				fmt.Println("Vous revenez au menu...")
			}

			if m.HP <=  0 {
				fmt.Println(m.Nom,"est vaincu !")
				finCombat(p, m, b, tour)
				return
			}
		case 3:
				fmt.Println(BROWN, "\n[--- Inventaire ---]", RESET)
			for j := range ii {
				if ii[j].Nombre > 0 && ii[j].Type != "" {
					fmt.Println(j+1, ")", ii[j].Nom, "x", ii[j].Nombre)
				}
			}
			fmt.Println("0) Retour")

			var choixItem int
			fmt.Scan(&choixItem)

			if choixItem == 0 {
				break
			}

			idx := choixItem - 1
			if idx < 0 || idx >= len(ii) || ii[idx].Nombre ==  0 {
				fmt.Println("Choix invalide.")
				break
			}

			ii[idx].Nombre--

			if ii[idx].Type == "soin" {
				p.stats.HPact += ii[idx].Effet
				if p.stats.HPact > p.stats.HPmax {
					p.stats.HPact = p.stats.HPmax
				}
				fmt.Println("Vous buvez", ii[idx].Nom, "! +", GREEN, ii[idx].Effet, "HP", RESET)
			} else if ii[idx].Type == "mana" {
				p.stats.MANA += ii[idx].Effet
				fmt.Println("Vous buvez", ii[idx].Nom, "! +", BLEUCLAIR, ii[idx].Effet, "MANA", RESET)
			} else if ii[idx].Type == "degats" {
				m.HP -= ii[idx].Effet
				fmt.Println("Vous lancez", ii[idx].Nom, "! -", GREEN, ii[idx].Effet, "HP", RESET, "au monstre")
			} else {
				fmt.Println("Cet objet ne peut pas être utilisé en combat.")
			}

			if m.HP <= 0 {
				fmt.Println(m.Nom, "est vaincu !")
				finCombat(p, m, b, tour)
				return
			}
		case 4:
			if m.Boss {
				fmt.Println(RED, "\n[!!! 🚧 IMPOSSIBLE DE FUIR LE GRAND LORD KODOÏD 🚧 !!!]", RESET)
			} else {
			fmt.Println("Vous tentez de fuir...")
			if rand.Intn(100) < 80 {
				fmt.Println("Vous avez réussi à fuir !")
				return
			} else {
				fmt.Println("Vous n'avez pas réussi à fuir !")
			}
		}
		default:
			fmt.Println("Si tu fais rien, tu vas te faire hagar salement 😈")
		}

		if m.HP > 0 {
			degats := m.Dgt
			armure := false
			for j := range ii {
				if ii[j].Nom == "Armure du damné" && ii[j].Nombre > 0 {
					armure = true
					break
		}
	}
	if armure {
		degats = int(float64(m.Dgt) * 0.85)
		fmt.Println("🛡️ Votre", GREY, "Armure du damné", RESET, "absorbe une partie des", RED, "dégâts", RESET, "!")
	}
	p.stats.HPact -= degats
	fmt.Println(m.Nom, "vous attaque et vous inflige",RED, degats, "dégâts", RESET, "!")
	fmt.Println("Il vous reste", GREEN, p.stats.HPact, "HP.", RESET)
}

		}

		if p.stats.HPact <=  0 {
			fmt.Println("\n Vous êtes", RED, "mort", RESET, "...Votre âme retourne au", GREY, "cimetière 💀", RESET)
			os.Exit(0)
		}
	}

func finCombat(p *Personnage, m Monstre, b *bell, tour int) {
	fmt.Println("Combat terminé en", tour, "tours !")
	fmt.Println("\nL'âme de", m.Nom,"se libère...")
	fmt.Println("1. Purifier", GREY, "l'âme", RESET, "(l'infuser dans ton arme)")
	fmt.Println("2. Sceller", GREY, "l'âme", RESET, "dans la", YELLOW, "bell", RESET)

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		switch p.classe {
		case "Vampire":
			p.stats.MANAmax +=  10
			fmt.Println("Ton arme absorbe", GREY, "l'âme",RESET, "!",BLEUCLAIR, "+10 MANA",RESET)
		case "Humain":
			p.stats.DGT +=  1
			fmt.Println("Ton arme absorbe",GREY, "l'âme",RESET, "!", RED, "+1 DGT",RESET)
		case "Slime":
			p.stats.HPmax +=  5
			p.stats.HPact +=  5
			fmt.Println("Ton arme absorbe",GREY, "l'âme",RESET, "!",GREEN, "+5 HP", RESET, "max")
		}
	} else {
		switch m.Nom {
		case "Spectre":
			b.Spectre += 1
		case "Possédé":
			b.Possede += 1
		case "Hybride":
			b.Hybride += 1
		}
		fmt.Println(GREY, "L'âme", RESET, "de", m.Nom, "est scellée dans ta", YELLOW, "bell", RESET, ".")
	}
}
