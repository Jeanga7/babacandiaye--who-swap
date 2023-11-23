package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// TO RUN ------>  |-->  go run tests/generate_nbrs.go 10 5 > tests/output.sh ; ./tests/output.sh

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
	if length > 200 {
		fmt.Println("Error: length should be less than or equal to 200.")
		return ""
	}
 
	rand.Seed(time.Now().UnixNano())
	numbers := rand.Perm(200)[:length]
 
	var lineBuilder strings.Builder
 
	for _, number := range numbers {
		lineBuilder.WriteString(strconv.Itoa(number + 1))
		lineBuilder.WriteString(" ")
	}
 
	return strings.TrimSpace(lineBuilder.String())
 }