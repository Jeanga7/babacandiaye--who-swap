package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// TO RUN ------>  |-->  `go run tests/generate_nbrs.go 10 3 > test/output.md`

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 2 {
		numLines, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erreur : veuillez spécifier un nombre valide de lignes.")
			return
		}

		numNumbers, err := strconv.Atoi(os.Args[2])
		if err != nil || numNumbers < 1 {
			fmt.Println("Erreur : veuillez spécifier un nombre valide de nombres par ligne (au moins 1).")
			return
		}

		for i := 0; i < numLines; i++ {
			line := generateRandomLine(numNumbers)
			fmt.Print("go run . " + strconv.Quote(line) + "\n")
		}
	} else {
		fmt.Println("Veuillez spécifier le nombre de lignes et le nombre de nombres par ligne en tant qu'arguments.")
	}
}

func generateRandomLine(length int) string {
	var lineBuilder strings.Builder

	for j := 0; j < length; j++ {
		lineBuilder.WriteString(strconv.Itoa(rand.Intn(5) + 1))
		if j < length-1 { // Ajouter un espace après chaque nombre, sauf après le dernier
			lineBuilder.WriteString(" ")
		}
	}

	return lineBuilder.String()
}