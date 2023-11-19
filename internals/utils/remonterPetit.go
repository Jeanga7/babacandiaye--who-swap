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
	if len(*pileA) < 2 {
		return
	}

	minIndex := 0
	minValue := (*pileA)[0].(int)

	for i, v := range *pileA {
		if v.(int) < minValue {
			minIndex = i
			minValue = v.(int)
		}
	}

	if minIndex == 1 && !EstTrie(pileA, pileB) {
		cmd.Sa(pileA)
		CompterMoves()
		fmt.Println("Après sa 5RP:", *pileA)
		return
	}

	moves := minIndex
	pivot := len(*pileA) / 2

	if minIndex > pivot {
		moves = len(*pileA) - minIndex
	}

	if minIndex < pivot && !EstTrie(pileA, pileB) {
	for i := 0; i < moves; i++ {
			cmd.Ra(pileA)
			CompterMoves()
			fmt.Println("Après ra 5RP:", *pileA)
		} 
	} else if minIndex >= pivot && !EstTrie(pileA, pileB){
		if pileA.Size()%2 == 0 {
	for i := 0; i < moves; i++ {
			cmd.Rra(pileA)
			CompterMoves()
			fmt.Println("Après rra 5RP:", *pileA)
		}
	}else{
		for i := 0; i < moves+1; i++ {
			cmd.Rra(pileA)
			CompterMoves()
			fmt.Println("Après rra 5RP:", *pileA)
		}
		}
	}
	if minIndex == 0 && !EstTrie(pileA, pileB) {
		cmd.Pb(pileA, pileB)
		CompterMoves()
		fmt.Println("Après pb 5:", *pileA, *pileB)
	} else {
		fmt.Println("Déjà triés avant pb 5:", *pileA, *pileB)
	}
}
