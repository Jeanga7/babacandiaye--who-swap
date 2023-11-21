package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

// cette fonction gere le trie des valeurs inférieurs à 6
func Less6Sorted(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, low []int, i int, n interface{}, pivot int) {
	if PileA.Size() == 4 {
		for _, l := range low {
			// maxA est au debut et un des low à la 2iem pos
			if (n == l && i == 1) && maxA == 0 {
				cmd.Ra(PileA)
				fmt.Println("Ra")
				models.Count++
				return
			} else if minA == 1 && maxA != 0 {
				// minA est à la 2iem pos
				cmd.Sa(PileA)
				fmt.Println("Sa")

				models.Count++
				return
			} else if (n == l && i == 0) {
				// low à la pos[0]
				cmd.Pb(PileA, PileB)
				fmt.Println("Pb")
				models.Count++
				return
			}
		}
		// minA est à la 2iem pos et maxA à la pos[0]
		if maxA == 0 && minA == 1 {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
			return
		}
		// minA est proche du bas
		if minA == len(*PileA)-1 || (maxA == 0 && minA == len(*PileA)-1) || (maxA == 1 && (minA == 2 || minA == len(*PileA)-1)) {
			cmd.Rra(PileA)
			fmt.Println("Rra")
			models.Count++
			return
		}
		// minA est à la 2iem pos et maxA à la pos[0]
		if minA == 2 && maxA == len(*PileA)-1 {
			cmd.Sa(PileA)
			fmt.Println("Sa")
			models.Count++
			return
		}
		// maxA == 0 et minA != de la dernier pos
		if maxA == 0 && minA != len(*PileA)-1 {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
			return
		}
	} else {
		for _, l := range low {
			if (n == l && i == 0) && minA != 1 {
				// low à la pos[0]
				cmd.Pb(PileA, PileB)
				fmt.Println("Pb")
				models.Count++
				return
			} else if ((n == l && i == 1) && (maxA != 0 || minA == 0)) || ((n == l && i == 0) && minA == 1) {
				// low à la deuxiem pos
				cmd.Sa(PileA)
				fmt.Println("Sa")
				models.Count++
				return
			} else if n == l && i == len(*PileA)-1 {
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				return
			}
		}
		// min est proche du bas
		if minA >= pivot {
			cmd.Rra(PileA)
			fmt.Println("Rra")
			models.Count++
			return
		}
		if minA < pivot {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
			return
		}
		if maxA == 0 {
			cmd.Ra(PileA)
			fmt.Println("Ra")
			models.Count++
			return
		}
		if minA == 1 && maxA != 0 {
			cmd.Sa(PileA)
			fmt.Println("Sa")
			models.Count++
			return
		}
	}
}
