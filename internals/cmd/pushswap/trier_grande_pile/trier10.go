package trier_grande_pile

import (
	"fmt"
	"push/internals/cmd"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
)

func TrierDixNombres(pileA, pileB *models.Stack) {
	iterations := len(*pileA) - 3

	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés 6:", pileA)
		return
	}

	for i := 0; i < iterations; i++ {
		utils.RemonterPlusPetit(pileA, pileB)
		if !utils.EstTrie(pileA, pileB) {
			cmd.Pb(pileA, pileB)
			utils.CounterMoves()
			fmt.Println("Après pb 6:", *pileA, *pileB)
		} else {
			fmt.Println("Déjà triés avant pb 6:", *pileA, *pileB)
		}
		psPetit.TrierCinqNombres(pileA, pileB)
	}

	psPetit.TrierCinqNombres(pileA, pileB)

	for len(*pileB) > 0 {
		cmd.Pa(pileB, pileA)
		utils.CounterMoves()
		fmt.Println("Après pa 6:", *pileA, *pileB)
		psPetit.TrierCinqNombres(pileA, pileB)
	}
}
