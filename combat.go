package main

import (
	"os"
	"fmt"
	"math/rand"
)

const RESETTT   = "\u001b[0m"
const REDD      = "\u001b[31m"
const YELLOWW   = "\u001b[33m"
const GREENN    = "\u001b[32m"

type Monstre struct {
	Nom string
	HP  int
	Dgt int
	Boss bool
}

func foret(p *Personnage, b *bell) {
	fmt.Println("\n")
	fmt.Println("[-Vous etes dans la forêt,faites attention rebroussez chemin avant de mourir!-]")
	fmt.Println("1.", REDD, "Se battre", RESETTT)
	fmt.Println("2.", YELLOWW, "Rebrousser chemin", RESETTT)

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		fmt.Println("\n")
		fmt.Println("[-Qui voulez-vous affronter?-]")
		fmt.Println("1.Spectre")
		fmt.Println("2.Possédé")
		fmt.Println("3.Hybride")

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
	fmt.Println("[Le DU-DU-DU-DUEL contre un", m.Nom, "commence!]")

	tour := 0 
	
	for p.stats.HPact > 0 && m.HP > 0 {
		tour++ 
		fmt.Println("\n ⚔️  |=== Tour de combat", tour," ===| ⚔️")
		fmt.Println("Vous:", p.stats.HPact, "HP |", m.Nom, ":", m.HP, "HP")
		fmt.Println(p.stats.MANA, "MANA restant")
		fmt.Println("1.Attaquer")
		fmt.Println("2.Résonances")
		fmt.Println("3.Inventaire")
		fmt.Println("4.Fuir")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			m.HP -= p.stats.DGT
			fmt.Println("Vous infligez", p.stats.DGT, "degats à", m.Nom,"!")
			if m.HP <= 0 {
				fmt.Println(m.Nom,"est vaincu !")
				finCombat(p, m, b, tour)
				return
			}
		case 2:
			fmt.Println("1. Coup de Poing (8 dégâts, 8 MANA)")
			fmt.Println("2. Boule de Feu (18 dégâts, 18 MANA)")
			fmt.Println("3. Dévotion Vampirique (vide tout ton MANA en dégâts, 150 MANA min)")
			fmt.Println("4. Retour")

			var choixComp int
			fmt.Scan(&choixComp)

			switch choixComp {
			case 1:
				if p.stats.MANA >=  8 {
					p.stats.MANA -=  8
					m.HP -=  8
					fmt.Println("Vous lancez Coup de Poing !  8 dégâts à", m.Nom,"!")
				} else {
					fmt.Println("Pas assez de MANA !")
				}
			case 2:
				if p.stats.MANA >=  18 {
					p.stats.MANA -=  18
					m.HP -= 18
					fmt.Println("Vous lancez Boule de Feu !  18 dégâts à", m.Nom,"!")
				} else {
					fmt.Println("Pas assez de MANA !")
				}
			case 3:
				if p.stats.MANA >=  150 {
					m.HP -= p.stats.MANA
					fmt.Println("Vous libérez toute votre énergie !", p.stats.MANA,"dégâts à", m.Nom,"!")
					p.stats.MANA = 0
				} else {
					fmt.Println("Il vous faut au moins 150 MANA pour cette résonance !")
				}
			case 4:
				fmt.Println("Vous revenez au menu...")
			}

			if m.HP <=  0 {
				fmt.Println(m.Nom,"est vaincu !")
				finCombat(p, m, b, tour)
				return
			}
		case 3:
				fmt.Println("\n[--- Inventaire ---]")
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
		fmt.Println("Vous buvez", ii[idx].Nom, "! +", ii[idx].Effet, "HP")
	} else if ii[idx].Type == "mana" {
		p.stats.MANA += ii[idx].Effet
		fmt.Println("Vous buvez", ii[idx].Nom, "! +", ii[idx].Effet, "MANA")
	} else if ii[idx].Type == "degats" {
		m.HP -= ii[idx].Effet
		fmt.Println("Vous lancez", ii[idx].Nom, "! -", ii[idx].Effet, "HP au monstre")
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
				fmt.Println("\n[!!! IMPOSSIBLE DE FUIR LE GRAND LORD KODÏD !!!]")
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
			fmt.Println("Si tu fais rien, tu vas te faire hagar salement😈")
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
		fmt.Println("🛡️ Votre Armure du damné absorbe une partie des dégâts !")
	}
	p.stats.HPact -= degats
	fmt.Println(m.Nom, "vous attaque et vous inflige", degats, "dégâts !")
	fmt.Println("Il vous reste", p.stats.HPact, "HP.")
}

		}

		if p.stats.HPact <=  0 {
			fmt.Println("\n Vous êtes mort...Votre âme retourne au cimetière")
			os.Exit(0)
		}
	}

func finCombat(p *Personnage, m Monstre, b *bell, tour int) {
	fmt.Println("Combat terminé en", tour, "tours !")
	fmt.Println("\nL'âme de", m.Nom,"se libère...")
	fmt.Println("1. Purifier l'âme (l'infuser dans ton arme)")
	fmt.Println("2. Sceller l'âme dans la bell")

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		switch p.classe {
		case "Vampire":
			p.stats.MANAmax +=  10
			fmt.Println("Ton arme absorbe l'âme ! +10 MANA")
		case "Humain":
			p.stats.DGT +=  1
			fmt.Println("Ton arme absorbe l'âme ! +1 DGT")
		case "Slime":
			p.stats.HPmax +=  5
			p.stats.HPact +=  5
			fmt.Println("Ton arme absorbe l'âme ! +5 HP max")
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
		fmt.Println("L'âme de", m.Nom, "est scellée dans ta bell.")
	}
}
