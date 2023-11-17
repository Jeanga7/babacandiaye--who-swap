package pushswap

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

var count int

// The function TrierPetitePile sorts a stack of numbers by either sorting three numbers or sorting
// five numbers.
func TrierPetitePile(pileA, pileB *models.Stack) {
	fmt.Println("État initial:", *pileA, *pileB)
	if len(*pileA) <= 3 {
		TrierTroisNombres(pileA, pileB)
	} else {
		TrierCinqNombres(pileA, pileB)
	}
	fmt.Println("État final:", *pileA, *pileB)
	fmt.Printf("vous avez fait %v de moves\n", count)
}

// The function `TrierTroisNombres` sorts three numbers in a stack.
func TrierTroisNombres(pileA, pileB *models.Stack) {
	if len(*pileA) == 2 {
		if (*pileA)[0].(int) > (*pileA)[1].(int) {
			cmd.Sa(pileA)
			count++
			fmt.Println("Après sa:", *pileA, *pileB)
		} else {
			fmt.Println("Déjà triés :", *pileA, *pileB)
		}
		return
	} else if len(*pileA) == 1 {
		fmt.Println("Pas besoin de trier", *pileA, *pileB)
		return
	}

	if (*pileA)[0].(int) > (*pileA)[1].(int) {
		if (*pileA)[1].(int) > (*pileA)[2].(int) {
			// Cas 4 (3 2 1)
			cmd.Sa(pileA)
			count++
			cmd.Rra(pileA)
			count++
			fmt.Println("Après sa et rra:", *pileA, *pileB)
		} else if (*pileA)[0].(int) > (*pileA)[2].(int) {
			// Cas 3 (3 1 2)
			cmd.Ra(pileA)
			count++
			fmt.Println("Après ra:", *pileA, *pileB)
		} else {
			// Cas 1 (2 1 3)
			cmd.Sa(pileA)
			count++
			fmt.Println("Après sa:", *pileA, *pileB)
		}
	} else {
		if (*pileA)[1].(int) > (*pileA)[2].(int) {
			if (*pileA)[2].(int) > (*pileA)[0].(int) {
				// Cas 5 (1 3 2)
				cmd.Sa(pileA)
				count++
				fmt.Println("Après sa :", *pileA, *pileB)
				cmd.Ra(pileA)
				count++
				fmt.Println("Après ra :", *pileA, *pileB)
			} else {
				// Cas 2 (2 3 1)
				cmd.Rra(pileA)
				count++
				fmt.Println("Après rra:", *pileA, *pileB)
			}
		} else {
			// Cas 6 (1 2 3)
			fmt.Println("Déjà triés :", *pileA, *pileB)
		}
	}
}


func TrierCinqNombres(pileA, pileB *models.Stack) {
	if !EstTrie(pileA) {
		// Ajouter la vérification de la meilleure option avant de trier
		MeilleurMove(pileA)
		TrierTroisNombres(pileA, pileB)
		return
	}
	iterations := len(*pileA) - 3
	for i := 0; i < iterations; i++ {
		RemonterPlusPetit(pileA)
		cmd.Pb(pileA, pileB)
		count++
		fmt.Println("Après pb:", *pileA, *pileB)
		TrierTroisNombres(pileA, pileB)
	}

	TrierTroisNombres(pileA, pileB)

	for i := 0; i < iterations; i++ {
		cmd.Pa(pileB, pileA)
		count++
		fmt.Println("Après pa:", *pileA, *pileB)
		TrierTroisNombres(pileA, pileB)
	}
}

func EstTrie(pileA *models.Stack) bool {
	for i := 0; i < len(*pileA)-1; i++ {
		if (*pileA)[i].(int) > (*pileA)[i+1].(int) {
			return false
		}
	}
	return true
}

func MeilleurMove(pileA *models.Stack) {
	minIndex := 0
	minValue := (*pileA)[0].(int)

	for i, v := range *pileA {
		if v.(int) < minValue {
			minIndex = i
			minValue = v.(int)
		}
	}

	movesUp := minIndex
	movesDown := len(*pileA) - minIndex

	if movesUp < movesDown {
		cmd.Ra(pileA)
		count++
		fmt.Println("Après ra:", *pileA)
	} else {
		cmd.Rra(pileA)
		count++
		fmt.Println("Après rra:", *pileA)
	}
}

func RemonterPlusPetit(pileA *models.Stack) {
	if len(*pileA) < 4 {
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

	moves := minIndex
	pivot := len(*pileA) / 2
	if minIndex > pivot {
		moves = len(*pileA) - minIndex
	}

	for i := 0; i < moves; i++ {
		if minIndex < pivot {
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