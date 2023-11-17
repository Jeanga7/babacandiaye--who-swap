package trier_grande_pile

import (
	"fmt"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
)

var count int

func TrierGrandePile(pileA, pileB *models.Stack) {
	fmt.Println("État initial:", pileA)

	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés :", *pileA)
		return // Return the count here
	}

	if len(*pileA) == 6 {
		TrierSixNombres(pileA, pileB)
	} else {
		utils.RemonterPlusPetit(pileA, pileB)
		psPetit.TrierCinqNombres(pileA, &models.Stack{})
		TrierGrandePile(pileA, pileB)
	}

	fmt.Println("État final:", *pileA)
	fmt.Printf("vous avez fait %v de moves\n", count)
}
