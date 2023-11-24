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
	//contentB := (*pileB)[0].(int)
	secondMinIndex := SecondMinIndex(*pileA, minIndex)
	//secondMinValue := (*pileA)[secondMinIndex].(int)
	movesMin := minIndex
	movesSecond := secondMinIndex

	if movesSecond < movesMin {
		minIndex, secondMinIndex = secondMinIndex, minIndex
		movesMin = movesSecond
	}

	pivot := len(*pileA) / 2

	if minIndex >= pivot {
		movesMin = len(*pileA) - minIndex
	}
	contentB := 0
	if pileB.Size() > 0 {
		contentB = (*pileB)[0].(int)
	}
	if (minIndex == 1 || secondMinIndex == 1) && !EstTrie(pileA, pileB) && contentB < minValue {
		if contentB < minValue && pileB.Size() > 0 {
			fmt.Println("continue")
		}else {
			cmd.Sa(pileA)
			CounterMoves()
			fmt.Println("Apres sa 5RP x:", *pileA)
			CmdPb(pileA, pileB)
			return
		}
	}

	minIndex, shouldReturn := BeforePivot(minIndex, pivot, pileA, pileB, movesMin, estTriee, secondMinIndex)
	if shouldReturn {
		return
	}
	shouldReturn1 := AfterPivot(minIndex, pivot, pileA, pileB, movesMin, secondMinIndex, estTriee)
	if shouldReturn1 {
		return
	}
	CmdPb(pileA, pileB)
}

func AfterPivot(minIndex int, pivot int, pileA *models.Stack, pileB *models.Stack, movesMin int, secondMinIndex int, estTriee bool) bool {
	if minIndex > pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < movesMin; i++ {
			minIndex = MinIndex(*pileA)
			secondMinIndex = SecondMinIndex(*pileA, minIndex)
			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP f:", *pileA)
				CmdPb(pileA, pileB)
				return true
			} else {
				minIndex = MinIndex(*pileA)
				secondMinIndex = SecondMinIndex(*pileA, minIndex)
				minValue := (*pileA)[minIndex].(int)
				contentB := 0
				if pileB.Size() > 0 {
					contentB = (*pileB)[0].(int)
				}

				if minIndex == 0 || secondMinIndex == 0 && contentB < minValue {
					CmdPb(pileA, pileB)
				} else if minIndex > secondMinIndex || minIndex < secondMinIndex {
					if pileA.Size() > 3 {
						if pileB.Size() > 0 {
							cmd.Rra(pileA)
							CounterMoves()
							fmt.Println("Apres rra 5RP g k:", *pileA)
						} else {
							cmd.Rra(pileA)
							CounterMoves()
							fmt.Println("Apres rra 5RP g q:", *pileA)
						}
					}
				} else {
					cmd.Rra(pileA)
					CounterMoves()
					fmt.Println("Apres rra 5RP z:", *pileA)
				}
			}
		}
	}
	return false
}

func BeforePivot(minIndex int, pivot int, pileA *models.Stack, pileB *models.Stack, movesMin int, estTriee bool, secondMinIndex int) (int, bool) {
	if minIndex <= pivot && !EstTrie(pileA, pileB) {
		for i := 0; i < movesMin; i++ {
			minIndex = MinIndex(*pileA)

			if minIndex == 1 && !EstTrie(pileA, pileB) && estTriee {
				cmd.Sa(pileA)
				CounterMoves()
				fmt.Println("Apres sa 5RP k:", *pileA)
				CmdPb(pileA, pileB)
				return 0, true
			} else {
				minIndex = MinIndex(*pileA)
				secondMinIndex = SecondMinIndex(*pileA, minIndex)
				minValue := (*pileA)[minIndex].(int)
				contentB := 0
				if pileB.Size() > 0 {
					contentB = (*pileB)[0].(int)
				}
				if minIndex == 0 || secondMinIndex == 0 && contentB < minValue {
					CmdPb(pileA, pileB)
				} else if minIndex > secondMinIndex || minIndex < secondMinIndex {
					if pileB.Size() > 0 {
						cmd.Rra(pileA)
						CounterMoves()
						fmt.Println("Apres rra 5RP g d:", *pileA)
					}
				} else if (minIndex > secondMinIndex || minIndex < secondMinIndex) && contentB < minValue {
					if pileB.Size() > 0 {
						cmd.Rra(pileA)
						CounterMoves()
						fmt.Println("Apres rra 5RP g d:", *pileA)
					}
				} else {
					cmd.Ra(pileA)
					CounterMoves()
					fmt.Println("Apres ra 5RP g:", *pileA)
				}
			}
		}
	}
	return minIndex, false
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
