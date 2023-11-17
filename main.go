package main

import (
	"os"
	"push/internals/cmd/pushswap"
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
			pushswap.TrierPetitePile(&pileA, &pileB)
		} else {
			utils.Throw("Nombre d'entiers incorrect" + utils.Log())
		}
	} else {
		utils.Throw("Besoin d'un argument" + utils.Log())
	}
}
