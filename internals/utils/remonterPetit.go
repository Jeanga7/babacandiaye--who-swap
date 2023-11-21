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
	estTriee := EstPartieTriee(pileA, 1, len(*pileA)-1)
	if len(*pileA) < 4 {
		return
	}

	minIndex := MinIndex(*pileA)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	movesMin := minIndex
	movesSecond := secondMinIndex

	if movesSecond < movesMin {
		minIndex, secondMinIndex = secondMinIndex, minIndex
	}

	pivot := len(*pileA) / 2

	if minIndex > pivot {
		movesMin = len(*pileA) - minIndex
	}

	if minIndex == 1 && !EstTrie(pileA, pileB) {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP:", *pileA)
		CmdPb(pileA, pileB)
		return
	}

	if minIndex <= pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < movesMin; i++ {
			minIndex = MinIndex(*pileA)
			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP:", *pileA)
				CmdPb(pileA, pileB)
				return
			} else {
				cmd.Ra(pileA)
				CounterMoves()
				fmt.Println("Après ra 5RP :", *pileA)
			}
		}

		CmdPb(pileA, pileB)
		return
	}
	if minIndex > pivot && !EstTrie(pileA, pileB) {
		minIndex = MinIndex(*pileA)
		for i := 0; i < movesMin; i++ {
			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP:", *pileA)
				CmdPb(pileA, pileB)
				return
			} else {
				cmd.Rra(pileA)
				CounterMoves()
				fmt.Println("Apres rra 5RP z:", *pileA)
			}
		}
		CmdPb(pileA, pileB)
		return
	}
}

//______________________________________________________________________________

func CmdPb(pileA *models.Stack, pileB *models.Stack) {
	minIndex := MinIndex(*pileA)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	if (minIndex == 0 || secondMinIndex == 0)&& !EstTrie(pileA, pileB) {
		cmd.Pb(pileA, pileB)
		CounterMoves()
		fmt.Println("Apres pb 5 :", *pileA, *pileB)
	} else {
		fmt.Println("Déjà tries before pb 5:", *pileA, *pileB)
	}
}
