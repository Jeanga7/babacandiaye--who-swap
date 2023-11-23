package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

// cette fonction permet de trier les plus grands nombres
func LesstenSorted(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, low []int, pivot int) {
	for i, n := range *PileA {
		if minA == 0 {
			cmd.Pb(PileA, PileB)
			fmt.Println("Pb")
			models.Count++
			return
		}
		for _, l := range low {
			if (n == l && i == 0) && minA != 1 {
				// low à la pos[0]
				cmd.Pb(PileA, PileB)
				fmt.Println("Pb")
				models.Count++
				return
			}
		}
		if minA == 1 && maxA != 0 {
			cmd.Sa(PileA)
			fmt.Println("Sa")
			models.Count++
			return
		}
		if minA >= pivot {
			cmd.Rra(PileA)
			fmt.Println("Rra")
			models.Count++
			return
		}
		if maxA == 0 && minA != len(*PileA)-1 {
			cmd.Ra(PileA)
			fmt.Println("test")
			fmt.Println("Ra")
			models.Count++
			return
		}
		if minA < pivot {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
			return
		}
	}
}
