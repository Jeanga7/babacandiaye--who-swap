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
	contentB := (*pileB)[0].(int)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	secondMinValue := (*pileA)[secondMinIndex].(int)
	movesMin := minIndex
	movesSecond := secondMinIndex

	if movesSecond < movesMin {
		minIndex, secondMinIndex = secondMinIndex, minIndex
		movesMin = movesSecond
	}
	

	pivot := len(*pileA) / 2

	if minIndex >= pivot  {
		movesMin = len(*pileA) - minIndex
	}

	if (minIndex == 1 || secondMinIndex == 1) && !EstTrie(pileA, pileB) {
		cmd.Sa(pileA)
		CounterMoves()
		fmt.Println("Apres sa 5RP x:", *pileA)
		CmdPb(pileA, pileB)
		return
	}

	if minIndex <= pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < movesMin; i++ {
			minIndex = MinIndex(*pileA)
			contentB = (*pileB)[0].(int)
			minValue = (*pileA)[minIndex].(int)
			if contentB > minValue && pileB.Size() > 0 {
				fmt.Println("hello")
			}
			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP k:", *pileA)
				CmdPb(pileA, pileB)
				return
			} else {
				minIndex = MinIndex(*pileA)
				secondMinIndex = SecondMinIndex(*pileA, minIndex)
				if minIndex == 0 || secondMinIndex == 0 {
					CmdPb(pileA, pileB)
				} else if minIndex > secondMinIndex || minIndex < secondMinIndex {
					cmd.Rra(pileA)
					CounterMoves()
					fmt.Println("Apres rra 5RP g:", *pileA)
					
				} else {
					fmt.Println("min", minIndex, minValue)
					fmt.Println("sec", secondMinIndex, secondMinValue)
					fmt.Println("piv", pivot)
					cmd.Ra(pileA)
					CounterMoves()
					fmt.Println("Apres ra 5RP g:", *pileA)
				}
			}
		}
	}
	if minIndex > pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < movesMin; i++ {
			minIndex = MinIndex(*pileA)
			secondMinIndex = SecondMinIndex(*pileA, minIndex)
			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP f:", *pileA)
				CmdPb(pileA, pileB)
				return
			} else {
				minIndex = MinIndex(*pileA)
				if minIndex == 0 || secondMinIndex == 0 {
					CmdPb(pileA, pileB)
				} else {
					cmd.Rra(pileA)
					CounterMoves()
					fmt.Println("Apres rra 5RP z:", *pileA)
				}
			}
		}
	}
	CmdPb(pileA, pileB)
}

//______________________________________________________________________________

func CmdPb(pileA, pileB *models.Stack) {

	if pileA.Size() < 4 {
		return
	}

	minIndex := MinIndex(*pileA)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	
	if minIndex == 0 || secondMinIndex == 0 {
		cmd.Pb(pileA, pileB)
		CounterMoves()
		fmt.Println("Après pb RP :", *pileA, *pileB)
	} else {
		fmt.Println("Déjà trié avant pb RP:", *pileA, *pileB)
	}
}
/* // __________________________________________________________________________________
// Contains vérifie si l'élément est présent dans la pile.
func ContainsNbr(pileB *models.Stack, element int) bool {
	for _, value := range *pileB {
		if value == element {
			return true
		}
	}
	return false
}
 */