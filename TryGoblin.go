package main

import "fmt"

func AccesTryGoblin(p *Personnage, b *bell) {
	fmt.Println("\n")
	fmt.Println("|==Vous êtes dans l'arene du TryGoblin==|")
	fmt.Println("1.S'entrainer sur le TryGoblin")
	fmt.Println("2.Retourner au village")

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
		t := Monstre{Nom: "TryGoblin", HP: 999999999, Dgt: 1}
		combat(p, t, b)
	}
}
