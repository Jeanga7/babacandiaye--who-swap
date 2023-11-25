package trier_grande_pile

import (
	"fmt"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
)


func TrierGrandePile(pileA, pileB *models.Stack) {
	fmt.Println("État initial:", pileA)

	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés :", *pileA)
		return 
	}

	if len(*pileA) <= 100 {
		TrierDixNombres(pileA, pileB)
	} else {
		utils.RemonterPlusPetit(pileA, pileB)
		psPetit.TrierCinqNombres(pileA, &models.Stack{})
		TrierGrandePile(pileA, pileB)
	}

	fmt.Println("État final:", *pileA)
}
