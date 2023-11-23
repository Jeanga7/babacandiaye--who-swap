package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

func Less100Sorted(PileA *models.Stack, PileB *models.Stack, low []int, pivot int) {
	for _, l := range low {
		for i, n := range *PileA {
			if n.(int) <= l && i == 0 {
				// pb
				cmd.Pb(PileA, PileB)
				fmt.Println("Pb")
				models.Count++
				return
			}
			if n.(int) <= l && i == 1 {
				// sa
				cmd.Sa(PileA)
				fmt.Println("Sa")
				models.Count++
				return
			}
			if (n.(int) < pivot) && (n.(int) <= l) {
				// ra
				cmd.Ra(PileA)
				fmt.Println("Ra")
				models.Count++
				return
			} else if (n.(int) >= pivot) && (n.(int) <= l) {
				// rra
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				return
			}
		}
	}
}

