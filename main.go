package main

import (
	"fmt"
	"os"
	psGrand "push/internals/cmd/pushswap/trier_grande_pile"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var pileA models.Stack
		var pileB models.Stack

		// Validation de l'argument
		if strings.Contains(arg, " ") {
			for _, ch := range arg {
				if (ch >= '0' && ch <= '9') || ch == 32 {
					continue
				} else {
					utils.Throw("Argument non valide" + utils.Log())
					return
				}
			}
		} else {
			utils.Throw("Argument non valide" + utils.Log())
			return
		}

		// Conversion de l'argument en entiers
		entiers := strings.Split(arg, " ")
		for _, v := range entiers {
			n, err := strconv.Atoi(v)
			if err != nil {
				utils.Throw("Erreur de conversion en entier" + utils.Log())
				return
			}
			pileA = append(pileA, n)
		}

		// Vérification du nombre d'entiers
		if len(pileA) <= 5 {
			psPetit.TrierPetitePile(&pileA, &pileB)
		} else if len(pileA) > 5 || len(pileA) <= 10 {
			psGrand.TrierGrandePile(&pileA, &pileB)
		} else {
			utils.Throw("Nombre d'entiers incorrect" + utils.Log())
			return
		}
	} else {
		utils.Throw("Besoin d'un argument" + utils.Log())
		return
	}
	fmt.Println("    ///////////////////////////////////////////")
	fmt.Println("   ///                                     ///")
	fmt.Printf("  ///  NOMBRE TOTAL DE MOUVEMENTS: %v ///\n", utils.Count)
	fmt.Println(" ///                                     ///")
	fmt.Println("///////////////////////////////////////////")
}
