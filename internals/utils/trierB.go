package utils

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)


func TrierB(pileA, pileB *models.Stack) {
	maxIndexA := MaxIndex(*pileA)
	minIndexA := MinIndex(*pileA)
	//maxIndexB := MaxIndex(*pileB)
	minIndexB := MinIndex(*pileB)
	if len(*pileA) < 2 {
		fmt.Println("Pas besoin de trier", *pileA, *pileB)
		return
	}
	if pileB.Size() == 2 {
		if minIndexB == 0 && maxIndexA == 1 && minIndexA == 0 {
			cmd.Ss(pileA, pileB)
			CounterMoves()
			fmt.Println("Après ss 3:", *pileA, *pileB)
		}else if minIndexB == 0 && maxIndexA == 0 && minIndexA == 2{

		} else if minIndexB == 0 {
			cmd.Sb(pileB)
			CounterMoves()
			fmt.Println("Après sb 3:", *pileA, *pileB)
		} else {
			EstTrie(pileA, pileB)
		}
		return
	}


	// if len(*pileA) == 3 {
	// 	if maxIndex == 0 && minIndex == pileB.Size()-1 {
	// 		cmd.Sa(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après sa 3 :", *pileA, *pileB)
	// 		cmd.Rra(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après rra 3 A:", *pileA, *pileB)
	// 	} else if maxIndex == 0 && minIndex == 1 {
	// 		cmd.Ra(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après ra 3:", *pileA, *pileB)
	// 	} else if maxIndex == 1 && minIndex == pileB.Size()-1 {
	// 		cmd.Rra(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après rra 3:", *pileA, *pileB)
	// 	} else if maxIndex == 1 && minIndex == 0 {
	// 		cmd.Sa(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après sa 3 :", *pileA, *pileB)
	// 		cmd.Ra(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après ra 3 A:", *pileA, *pileB)
	// 	} else if maxIndex == pileB.Size()-1 && minIndex == 1 {
	// 		cmd.Sa(pileA)
	// 		CounterMoves()
	// 		fmt.Println("Après sa 3:", *pileA, *pileB)
	// 	}
	// }
	EstTrie(pileA, pileB)
}
