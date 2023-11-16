package pushswap

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

func TrierTroisNombres(pileA, pileB *models.Stack) {
	fmt.Println("État initial:", *pileA, *pileB)

	if len(*pileA) == 2 {
		if (*pileA)[0].(int) > (*pileA)[1].(int) {
			cmd.Sa(pileA)
			fmt.Println("Après sa:", *pileA, *pileB)
		} else {
			fmt.Println("Déjà triés :", *pileA, *pileB)
		}
		return
	} else if len(*pileA) == 1 {
		fmt.Println("Pas besoin de trier", *pileA, *pileB)
		return
	}

	if (*pileA)[0].(int) > (*pileA)[1].(int) {
		if (*pileA)[1].(int) > (*pileA)[2].(int) {
			// Cas 4 (3 2 1)
			cmd.Sa(pileA)
			cmd.Rra(pileA)
			fmt.Println("Après sa et rra:", *pileA, *pileB)
		} else if (*pileA)[0].(int) > (*pileA)[2].(int) {
			// Cas 3 (3 1 2)
			cmd.Ra(pileA)
			fmt.Println("Après ra:", *pileA, *pileB)
		} else {
			// Cas 1 (2 1 3)
			cmd.Sa(pileA)
			fmt.Println("Après sa:", *pileA, *pileB)
		}
	} else {
		if (*pileA)[1].(int) > (*pileA)[2].(int) {
			if (*pileA)[2].(int) > (*pileA)[0].(int) {
				// Cas 5 (1 3 2)
				cmd.Sa(pileA)
				cmd.Ra(pileA)
				fmt.Println("Après sa et ra:", *pileA, *pileB)
			} else {
				// Cas 2 (2 3 1)
				cmd.Rra(pileA)
				fmt.Println("Après rra:", *pileA, *pileB)
			}
		} else {
			// Cas 6 (1 2 3)
			fmt.Println("Déjà triés :", *pileA, *pileB)
		}
	}
}
