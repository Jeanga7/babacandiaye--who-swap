package trierpetitpile

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)



/*___________________________________________________The function "TrierCinqNombres" sorts two stacks
                                                     of numbers by using the "RemonterPlusPetit" and
                                                     "TrierTroisNombres" functions.*/
func TrierCinqNombres(pileA, pileB *models.Stack) {
	if EstTrie(pileA) {
		fmt.Println("Déjà triés :", *pileA, *pileB)
		return
	}

	iterations := len(*pileA) - 3

	for i := 0; i < iterations; i++ {
		RemonterPlusPetit(pileA)
		if !EstTrie(pileA) {
			cmd.Pb(pileA, pileB)
			count++
			fmt.Println("Après pb:", *pileA, *pileB)
		} else {
			fmt.Println("Déjà triés avant pb :", *pileA, *pileB)
		}
		TrierTroisNombres(pileA, pileB)
	}

	TrierTroisNombres(pileA, pileB)

	for len(*pileB) > 0 {
		cmd.Pa(pileB, pileA)
		count++
		fmt.Println("Après pa:", *pileA, *pileB)
		TrierTroisNombres(pileA, pileB)
	}
}

/*_______________________________________________The function "EstTrie" checks 
                                                 if a stack is sorted in ascending order.*/
func EstTrie(pileA *models.Stack) bool {
	for i := 0; i < len(*pileA)-1; i++ {
		if (*pileA)[i].(int) > (*pileA)[i+1].(int) {
			return false
		}
	}
	return true
}

/*_______________________________________________The function "RemonterPlusPetit" is used to move
                                                  the smallest element in a stack to the top. */
func RemonterPlusPetit(pileA *models.Stack) {
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

	if minIndex == 1 && !EstTrie(pileA) {
		cmd.Sa(pileA)
		count++
		fmt.Println("Après sa:", *pileA)
		return
	}

	moves := minIndex
	pivot := len(*pileA) / 2

	if minIndex > pivot {
		moves = len(*pileA) - minIndex
	}

	for i := 0; i < moves; i++ {
		if minIndex < pivot && !EstTrie(pileA) {
			cmd.Ra(pileA)
			count++
			fmt.Println("Après ra:", *pileA)
		} else {
			cmd.Rra(pileA)
			count++
			fmt.Println("Après rra:", *pileA)
		}
	}
}
