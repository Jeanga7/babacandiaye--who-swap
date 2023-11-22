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
	minValue := (*pileA)[minIndex].(int)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	secondMinValue := (*pileA)[secondMinIndex].(int)
	movesMin := minIndex
	movesSecond := secondMinIndex

	if movesSecond < movesMin {
		minIndex, secondMinIndex = secondMinIndex, minIndex
		movesMin = movesSecond
	}

	pivot := len(*pileA) / 2

	if minIndex > pivot {
		movesMin = len(*pileA) - minIndex
	}

	if (minIndex == 1 || secondMinIndex == 1) && !EstTrie(pileA, pileB) {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP x:", *pileA)
		CmdPb(pileA, pileB)
		return
	}

	// Vérifier si minValue est dans pileB et remonter secondMinValue si nécessaire
	if ContainsNbr(pileB, minValue) && !EstTrie(pileA, pileB) && estTriee {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP k:", *pileA)
		CmdPb(pileA, pileB)
		return
	}

	// Vérifier si secondMinValue est dans pileB et remonter minValue si nécessaire
	if ContainsNbr(pileB, secondMinValue) && !EstTrie(pileA, pileB) && estTriee {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP k:", *pileA)
		CmdPb(pileA, pileB)
		return
	}

	for i := 0; i < movesMin; i++ {
		minIndex = MinIndex(*pileA)
		if minIndex == 0 || secondMinIndex == 0 {
			CmdPb(pileA, pileB)
		} else {
			cmd.Ra(pileA)
			CounterMoves()
			fmt.Println("Apres ra 5RP:", *pileA)
		}
	}

	CmdPb(pileA, pileB)
}

//______________________________________________________________________________

func CmdPb(pileA, pileB *models.Stack) {
	minIndex := MinIndex(*pileA)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	if (minIndex == 0 || secondMinIndex == 0) && !EstTrie(pileA, pileB) && pileA.Size() > 3 {
		cmd.Pb(pileA, pileB)
		CounterMoves()
		fmt.Println("Apres pb RP :", *pileA, *pileB)
	} else {
		fmt.Println("Déjà tries before pb RP:", *pileA, *pileB)
	}
}

// __________________________________________________________________________________
// Contains vérifie si l'élément est présent dans la pile.
func ContainsNbr(pileB *models.Stack, element int) bool {
	for _, value := range *pileB {
		if value == element {
			return true
		}
	}
	return false
}
