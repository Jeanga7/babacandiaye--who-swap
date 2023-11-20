package utils

import (
	"fmt"

	"push/internals/cmd"
	"push/internals/models"
)

/*
_______________________________________________The function "RemonterPlusPetit" is used to move

	the smallest element in a stack to the top.*/

func RemonterPlusPetit(pileA, pileB *models.Stack) {
	if len(*pileA) < 4 {
		return
	}

	minIndex := MinIndex(*pileA)
	moves := minIndex
	pivot := len(*pileA) / 2
	if minIndex > pivot {
		moves = len(*pileA) - minIndex
	}

	if minIndex == 1 && !EstTrie(pileA, pileB) {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP:", *pileA)
		return
	}

	if minIndex <= pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < moves; i++ {
			cmd.Ra(pileA)
			CounterMoves()
			fmt.Println("Après ra 5RP :", *pileA)
		}
	}
	if minIndex > pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < moves; i++ {
			cmd.Rra(pileA)
			CounterMoves()
			fmt.Println("Apres rra 5RP z:", *pileA)
		}

	}

	if minIndex == 0 && !EstTrie(pileA, pileB) {
		cmd.Pb(pileA, pileB)
		CounterMoves()
		fmt.Println("Apres pb 5 :", *pileA, *pileB)
	} else {
		fmt.Println("Déjà tries before pb 5:", *pileA, *pileB)
	}

}
