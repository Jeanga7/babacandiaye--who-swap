package trierpetitpile

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
	"push/internals/utils"
)

/*
________________________________________________________The function "TrierTroisNombres"
                                                         sorts three numbers in a stack.*/
func TrierTroisNombres(pileA, pileB *models.Stack) {
	maxIndex := utils.MaxIndex(*pileA)
	minIndex := utils.MinIndex(*pileA)
	if len(*pileA) < 2 {
		fmt.Println("Pas besoin de trier", *pileA, *pileB)
		return
	}

	if len(*pileA) == 2 {
		if maxIndex == 0{
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Après sa 3:", *pileA, *pileB)
		} else {
			utils.EstTrie(pileA, pileB)
		}
		return
	}

	if len(*pileA) == 3 {
		if maxIndex == 0 && minIndex == 2 {
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Après sa 3 :", *pileA, *pileB)
			cmd.Rra(pileA)
			utils.CounterMoves()
			fmt.Println("Après rra 3 A:", *pileA, *pileB)
		}else if maxIndex == 0 && minIndex == 1{
			cmd.Ra(pileA)
			utils.CounterMoves()
			fmt.Println("Après ra 3:", *pileA, *pileB)
		} else if maxIndex == 1 && minIndex == 2{
			cmd.Rra(pileA)
			utils.CounterMoves()
			fmt.Println("Après rra 3:", *pileA, *pileB)
		} else if maxIndex == 1 && minIndex == 0{
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Après sa 3 :", *pileA, *pileB)
			cmd.Ra(pileA)
			utils.CounterMoves()
			fmt.Println("Après ra 3 A:", *pileA, *pileB)
		} else if maxIndex == 2 && minIndex == 1{
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Après sa 3 :", *pileA, *pileB)
		}
	}
	utils.EstTrie(pileA, pileB)
}
