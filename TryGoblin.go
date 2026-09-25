package main

import "fmt"

const BBLANC   = "\u001b[37m"
const ROSEE    = "\033[38;5;201m"

func AccesTryGoblin(p *Personnage, b *bell) {
	fmt.Println("\n")
	fmt.Println(ROSEE, "|==Vous êtes dans l'arene du TryGoblin==|", BBLANC)
	fmt.Println("\n")
	fmt.Println("1.S'entrainer sur le", ROSEE, "TryGoblin",BBLANC)
	fmt.Println("2.Retourner au village")

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		t := Monstre{Nom: "TryGoblin", HP: 999999999, Dgt: 1}
		combat(p, t, b)
	}
}
