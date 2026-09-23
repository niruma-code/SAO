package main

import "fmt"

type bell struct {
	Spectre int
	Possede int
	Hybride int
}

type inventory struct {
	Nom   string
	Type string
	Effet int
	Nombre int
}


func afficherinventaire(i *[]inventory) {
	fmt.Println("\n")
	fmt.Println("   /===Inventaire===/   ")
	for j := range *i {
		if (*i)[j].Nombre > 0 {
			fmt.Println((*i)[j].Nom, ":", (*i)[j].Nombre)
		}
	}
	fmt.Println("\n") 
}