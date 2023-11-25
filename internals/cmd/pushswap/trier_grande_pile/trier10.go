package trier_grande_pile

import (
	"fmt"
	"push/internals/cmd"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
)

func TrierDixNombres(pileA, pileB *models.Stack) {
	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés 6:", *pileA, *pileB)
		return
	}
	
	utils.RemonterPlusPetit(pileA, pileB)


	if pileB.Size() <= 5 {
		fmt.Println("debut trie a 5")
		utils.RadixSort(pileA)
		fmt.Println("fin trie a 5")
	}

	for len(*pileB) > 0 {
		if (*pileA)[0].(int) > (*pileA)[1].(int) {
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Apres sa 6:", *pileA, *pileB)
		}
		cmd.Pa(pileB, pileA)
		utils.CounterMoves()
		fmt.Println("Après pa 6:", *pileA, *pileB)
		psPetit.TrierTroisNombres(pileA, pileB)
	}

	if !utils.EstTrie(pileA, pileB) && pileB.Size() == 0 {
		cmd.Sa(pileA)
		utils.CounterMoves()
		fmt.Println("Apres sa 6:", *pileA, *pileB)
	}
}
