package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

// cette fonction gere le trie des petites valeurs ie inferieur ou égale à 3
func Less3Sorted(PileA *models.Stack, PileB *models.Stack, minA int) {
	if PileA.Size() == 2 {
		cmd.Sa(PileA)
		fmt.Println("Sa")
		models.Count++
	} else if PileA.Size() == 3 {
		// ____________
		//        2 1 3
		if minA == 1 && (*PileA)[2].(int) > (*PileA)[0].(int) {
			cmd.Sa(PileA)
			fmt.Println("Sa")
			models.Count++
		}
		// _____________________
		//        3 1 2 || 3 2 1
		if ((minA == 2 && (*PileA)[0].(int) > (*PileA)[1].(int)) && IsinvSorted(PileA)) || (minA == 1 && (*PileA)[0].(int) > (*PileA)[2].(int)) {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
		}
		// _______________________
		//        2 3 1  ou  1 3 2
		if (minA == 2 && (*PileA)[1].(int) > (*PileA)[2].(int)) || (minA == 0 && (*PileA)[1].(int) > (*PileA)[2].(int)) {
			cmd.Rra(PileA)
			fmt.Println("Rra")
			models.Count++
		}
	}
}
