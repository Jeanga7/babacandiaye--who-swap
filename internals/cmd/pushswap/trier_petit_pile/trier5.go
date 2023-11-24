package trierpetitpile

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
	"push/internals/utils"
)

/*
___________________________________________________The function "TrierCinqNombres" sorts two stacks
                                                    of numbers by using the "RemonterPlusPetit" and
	                                                "TrierTroisNombres" functions.*/

func TrierCinqNombres(pileA, pileB *models.Stack) {
	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés 5:", *pileA, *pileB)
		return
	}

	estTriee := utils.EstPartieTriee(pileA, 1, len(*pileA)-1)
	maxIndex := utils.MaxIndex(*pileA)
	minIndex := utils.MinIndex(*pileA)

	if estTriee && maxIndex == 0 && pileA.Size() > 3 {
		cmd.Ra(pileA)
		utils.CounterMoves()
		fmt.Println("Après ra 5:", *pileA, *pileB)
		return
	}

	if estTriee && minIndex == 0 && pileA.Size() > 3 {
		cmd.Rra(pileA)
		utils.CounterMoves()
		fmt.Println("Apres rra 5:", *pileA, *pileB)
		return
	}

	iterations := len(*pileA) - 3

	for i := 0; i < iterations; i++ {
		utils.RemonterPlusPetit(pileA, pileB)
	}

	utils.TrierB(pileA, pileB)
	TrierTroisNombres(pileA, pileB)

	for len(*pileB) > 0 {
		if (*pileA)[0].(int) > (*pileA)[1].(int) {
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Apres sa 5:", *pileA, *pileB)
		}
		cmd.Pa(pileB, pileA)
		utils.CounterMoves()
		fmt.Println("Après pa 5:", *pileA, *pileB)
		TrierTroisNombres(pileA, pileB)
	}

	if !utils.EstTrie(pileA, pileB) && pileB.Size() == 0 {
		cmd.Sa(pileA)
		utils.CounterMoves()
		fmt.Println("Apres sa 5:", *pileA, *pileB)
	}
}
