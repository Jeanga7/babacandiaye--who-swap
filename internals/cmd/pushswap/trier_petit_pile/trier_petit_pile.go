package trierpetitpile

import (
	"fmt"
	"push/internals/models"
)
var count int
/*________________________________________________The function "TrierPetitePile" sorts two stacks 
                                                   of numbers by using other helper functions. */
func TrierPetitePile(pileA, pileB *models.Stack) {
	fmt.Println("État initial:", *pileA, *pileB)

	if EstTrie(pileA) {
		fmt.Println("Déjà triés :", *pileA, *pileB)
		return
	}

	if len(*pileA) <= 3 {
		TrierTroisNombres(pileA, pileB)
	} else if len(*pileA) > 3 || len(*pileA) <= 5{
		TrierCinqNombres(pileA, pileB)
	}

	fmt.Println("État final:", *pileA, *pileB)
	fmt.Printf("vous avez fait %v de moves\n", count)
}
