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
			// } else if ((n == l && i == 1) && (maxA != 0 || minA == 0)) || ((n == l && i == 0) && minA == 1) {
			// 	// low à la deuxiem pos
			// 	cmd.Sa(PileA)
			// 	fmt.Println("Sa")
			// 	models.Count++
			// 	return
			// }
		}
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
}

// brouillon
// for i, n := range *PileA {
// 	if minA == 0 {
// 		cmd.Pb(PileA, PileB)
// 		fmt.Println("Pb")
// 		models.Count++
// 		return
// 	}
// 	for _, l := range low {
// 		if (n == l && i == 0) && minA != 1 {
// 			// low à la pos[0]
// 			cmd.Pb(PileA, PileB)
// 			fmt.Println("Pb")
// 			models.Count++
// 			return
// 		}
// 	}
// 	// 	} else if ((n == l && i == 1) && (maxA != 0 || minA == 0)) || ((n == l && i == 0) && minA == 1) {
// 	// 		// low à la deuxiem pos
// 	// 		cmd.Sa(PileA)
// 	// 		fmt.Println("Sa")
// 	// 		models.Count++
// 	// 		return
// 	// 	} else if n == l && i == len(*PileA)-1 {
// 	// 		cmd.Rra(PileA)
// 	// 		fmt.Println("Rra")
// 	// 		models.Count++
// 	// 		return
// 	// 	} else if maxA == 0 && n == l && i == len(*PileA)-1 {
// 	// 		cmd.Rra(PileA)
// 	// 		fmt.Println("Rra")
// 	// 		models.Count++
// 	// 		return
// 	// 	}
// 	// }
// 	if minA == 1 && maxA != 0 {
// 		cmd.Sa(PileA)
// 		fmt.Println("Sa")
// 		models.Count++
// 		return
// 	}
// 	if minA >= pivot {
// 		cmd.Rra(PileA)
// 		fmt.Println("Rra")
// 		models.Count++
// 		return
// 	}
// 	if maxA == 0 && minA != len(*PileA)-1 {
// 		cmd.Ra(PileA)
// 		fmt.Println("test")
// 		fmt.Println("Ra")
// 		models.Count++
// 		return
// 	}
// 	if minA < pivot {
// 		cmd.Ra(PileA)
// 		fmt.Println("Ra")
// 		models.Count++
// 		return
// 	}
// }

// .............................................brouillon 2
// for i, n := range *PileA {
// 	for _, l := range low {
// 		if ((n == l && i == 0) && minA != 1) || minA == 0 {
// 			cmd.Pb(PileA, PileB)
// 			fmt.Println("Pb")
// 			models.Count++
// 			return
// 		}
// 		if ((n == l && i == 1) && (maxA != 0 && minA != 0)) || (minA == 1 && maxA != 0) {
// 			cmd.Sa(PileA)
// 			fmt.Println("Sa")
// 			models.Count++
// 			return
// 		}
// 		if n == l && i == len(*PileA)-1 {
// 			cmd.Rra(PileA)
// 			fmt.Println("Rra")
// 			fmt.Println("int")
// 			models.Count++
// 			return
// 		}
// 		// ........................
// 		if n == l && i == PileA.Size()-1 {
// 			cmd.Rra(PileA)
// 			fmt.Println("rra")
// 			models.Count++
// 			return
// 		}
// 		if n == l && i == 1 {
// 			cmd.Sa(PileA)
// 			fmt.Println("sa")
// 			// fmt.Println("je vais MAMa")
// 			models.Count++
// 			return
// 		}
// 	}

// 	//.........................
// 	if PileA.Size()%2 == 0 {
// 		if n == minA && i > PileA.Size()/2 {
// 			cmd.Rra(PileA)
// 			fmt.Println("rra")
// 			models.Count++
// 			return
// 		}
// 		if n == minA && i <= PileA.Size()/2 && i != 0 {
// 			cmd.Ra(PileA)
// 			fmt.Println("ra")
// 			models.Count++
// 			return
// 		}
// 	}
// 	if PileA.Size()%2 == 1 {
// 		if n == minA && i > PileA.Size()/2 && i != 0 {
// 			cmd.Rra(PileA)
// 			fmt.Println("rra")
// 			models.Count++
// 			return
// 		}
// 		if n == minA && i <= PileA.Size()/2 && i != 0 {
// 			cmd.Ra(PileA)
// 			fmt.Println("ra")
// 			models.Count++
// 			return
// 		}
// 	}
// 	if n == minA && i == 1 {
// 		cmd.Sa(PileA)
// 		fmt.Println("sa")
// 		models.Count++
// 		return
// 	}
// 	if n == minA && i > PileA.Size()/2 {
// 		cmd.Rra(PileA)
// 		fmt.Println("rra")
// 		models.Count++
// 		return
// 	}
// }
