package main

import "fmt"

const BOLD    = "\u001b[1m"
const GOLD    = "\u001b[33m"   
const WHITE   = "\u001b[37m"   
const MAGENTA = "\u001b[35m"
const CYAN    = "\u001b[36m"
const BLUE    = "\u001b[34m"
const BLINK   = "\u001b[5m"

func WhoIsBro() {
    quiSontIls()
}

func quiSontIls() {
    fmt.Println(GOLD + BOLD + "=== LÉGENDE ===" + RESET)
    fmt.Println(WHITE + "Le chef du village d'exorcistes saint du nord raconte qu'un jour," + RESET)
    fmt.Println(WHITE + "un exorciste possédé par le péché de la paresse atteindra un stade" + RESET)
    fmt.Println(WHITE + "où il invoquera un monstre si puissant qu'une préparation hors pair" + RESET)
    fmt.Println(WHITE + "sera assez pour avoir juste une chance." + RESET)

    fmt.Println()
    fmt.Println(BOLD + "Créateurs du jeu :" + RESET)

    fmt.Println(BLINK + RED + BOLD + "- LOAN" + RESET + "  " + CYAN + "⚡" + RESET)
    fmt.Println(BOLD + CYAN + "- AYRON" + RESET + " " + BLUE + "💎" + RESET)
    fmt.Println(BOLD + MAGENTA + "- ELEA" + RESET + "  " + GREEN + "🌸" + RESET)
    fmt.Println(BOLD + GREEN + "- NIKA" + RESET + "  " + GREEN + "🍀" + RESET)
}