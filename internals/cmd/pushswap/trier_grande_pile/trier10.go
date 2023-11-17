package trier_grande_pile

import (
	"fmt"
	"push/internals/cmd"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
)

func TrierSixNombres(pileA, pileB *models.Stack) {
	iterations := len(*pileA) - 3

	if psPetit.EstTrie(pileA) {
		fmt.Println("Déjà triés :", pileA)
		return
	}

	for i := 0; i < iterations; i++ {
		psPetit.RemonterPlusPetit(pileA)
		if !psPetit.EstTrie(pileA) {
			cmd.Pb(pileA, pileB)
			count++
			fmt.Println("Après pb:", *pileA, *pileB)
		} else {
			fmt.Println("Déjà triés avant pb :", *pileA, *pileB)
		}
		psPetit.TrierCinqNombres(pileA, pileB)
	}

	psPetit.TrierCinqNombres(pileA, pileB)

	for len(*pileB) > 0 {
		cmd.Pa(pileB, pileA)
		count++
		fmt.Println("Après pa:", *pileA, *pileB)
		psPetit.TrierCinqNombres(pileA, pileB)
	}
}
