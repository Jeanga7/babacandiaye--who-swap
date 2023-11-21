package utils

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)


func TrierB(pileA, pileB *models.Stack) {
	maxIndex := MaxIndex(*pileA)
	minIndex := MinIndex(*pileA)
	if len(*pileA) < 2 {
		fmt.Println("Pas besoin de trier", *pileA, *pileB)
		return
	}
	if len(*pileA) == 2 {
		if maxIndex == 0{
			cmd.Sa(pileA)
			CounterMoves()
			fmt.Println("Après sa 3:", *pileA, *pileB)
		} else {
			EstTrie(pileA, pileB)
		}
		return
	}


	if len(*pileA) == 3 {
		if maxIndex == 0 && minIndex == pileB.Size()-1 {
			cmd.Sa(pileA)
			CounterMoves()
			fmt.Println("Après sa 3 :", *pileA, *pileB)
			cmd.Rra(pileA)
			CounterMoves()
			fmt.Println("Après rra 3 A:", *pileA, *pileB)
		} else if maxIndex == 0 && minIndex == 1 {
			cmd.Ra(pileA)
			CounterMoves()
			fmt.Println("Après ra 3:", *pileA, *pileB)
		} else if maxIndex == 1 && minIndex == pileB.Size()-1 {
			cmd.Rra(pileA)
			CounterMoves()
			fmt.Println("Après rra 3:", *pileA, *pileB)
		} else if maxIndex == 1 && minIndex == 0 {
			cmd.Sa(pileA)
			CounterMoves()
			fmt.Println("Après sa 3 :", *pileA, *pileB)
			cmd.Ra(pileA)
			CounterMoves()
			fmt.Println("Après ra 3 A:", *pileA, *pileB)
		} else if maxIndex == pileB.Size()-1 && minIndex == 1 {
			cmd.Sa(pileA)
			CounterMoves()
			fmt.Println("Après sa 3:", *pileA, *pileB)
		}
	}
	EstTrie(pileA, pileB)
}
