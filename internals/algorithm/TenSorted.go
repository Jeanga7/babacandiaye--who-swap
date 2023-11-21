package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

func LesstenSorted(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, low []int, i int, n interface{}, pivot int) {
	if minA == 0 {
		cmd.Pb(PileA, PileB)
		fmt.Println("Pb")
		models.Count++
		return
	}
	// if low == 0 pb
	// for _, l := range low {
	// 	if (n == l && i == 0) && minA != 1 {
	// 		// low à la pos[0]
	// 		cmd.Pb(PileA, PileB)
	// 		fmt.Println("Pb")
	// 		models.Count++
	// 		return
	// 	}
	// }
	// 	} else if ((n == l && i == 1) && (maxA != 0 || minA == 0)) || ((n == l && i == 0) && minA == 1) {
	// 		// low à la deuxiem pos
	// 		cmd.Sa(PileA)
	// 		fmt.Println("Sa")
	// 		models.Count++
	// 		return
	// 	} else if n == l && i == len(*PileA)-1 {
	// 		cmd.Rra(PileA)
	// 		fmt.Println("Rra")
	// 		models.Count++
	// 		return
	// 	} else if maxA == 0 && n == l && i == len(*PileA)-1 {
	// 		cmd.Rra(PileA)
	// 		fmt.Println("Rra")
	// 		models.Count++
	// 		return
	// 	}
	// }
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
