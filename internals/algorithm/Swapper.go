package algorithm

import (
	"fmt"
	"os"
	"push/internals/cmd"
	"push/internals/models"
	"runtime"
	"sort"
)

type Positions struct {
	Pos []int
}

func Throw(message string) {
	fmt.Printf("\033[38;2;255;0;0m%v\033[0m\n", message)
	os.Exit(1)
}

func Log() string {
	pc, _, line, _ := runtime.Caller(1)
	callinFunc := runtime.FuncForPC(pc).Name()
	errMessage := fmt.Sprintf(" %v at line %v", callinFunc, line)
	return errMessage
}

// permet de recuperer l'index des plus petites valeurs (low)
func CheckPos(a *models.Stack, low []int, pos *Positions) {
	pos.Pos = nil
	for i, n := range *a {
		for _, l := range low {
			if n.(int) == l {
				pos.Pos = append(pos.Pos, i)
				break
			}
		}
	}
}

// permet d'extraire les plus petites valeurs
func SmallestFromStack(PileA *models.Stack) []int {
	numToExtract := PileA.Size() / 2
	if numToExtract < 0 {
		numToExtract = 0
	}

	// Convertir la pile en une tranche pour le tri
	numbers := make([]int, PileA.Size())
	for i, v := range *PileA {
		// Assertion de type
		if val, ok := v.(int); ok {
			numbers[i] = val
		}
	}

	sort.Ints(numbers)
	smallestNumbers := numbers[:numToExtract]
	return smallestNumbers
}

// verifie si une pile est ascendant
func IsSorted(stack *models.Stack) bool {
	for i := 0; i < stack.Size()-1; i++ {
		if (*stack)[i].(int) > (*stack)[i+1].(int) {
			return false
		}
	}
	return true
}

// permet de swap les deux premier élément de la pile si le premier est superieur au deuxieme
func SortedA(stack *models.Stack) {
	if (*stack)[0].(int) > (*stack)[1].(int) {
		cmd.Sa(stack)
	}
}

// verifie si une pile est descendant
func IsinvSorted(stack *models.Stack) bool {
	for i := 0; i < stack.Size()-1; i++ {
		if (*stack)[i].(int) < (*stack)[i+1].(int) {
			return false
		}
	}
	return true
}

// permet d'extraire la pos de min du  pile A
func MinPile(a *models.Stack) int {
	min := 1000000000000
	pos := 0
	for i, x := range *a {
		if x.(int) < min {
			min = x.(int)
			pos = i
		}
	}
	return pos
}

// permet d'extraire la pos de max du pile A
func MaxPile(a *models.Stack) int {
	max := -1000000000000
	pos := 0
	for i, x := range *a {
		if x.(int) > max {
			max = x.(int)
			pos = i
		}
	}
	return pos
}

// code propre
func Swapper(PileA *models.Stack, PileB *models.Stack) {
	// var P Positions
outerloop:
	for !IsSorted(PileA) {
		for i, n := range *PileA {
			// CheckPos(PileA, lows, &P)
			low := SmallestFromStack(PileA)
			minA := MinPile(PileA)
			maxA := MaxPile(PileA)
			minB := MinPile(PileB)
			maxB := MaxPile(PileB)
			pivot := 0
			if PileA.Size()%2 == 0 {
				pivot = PileA.Size() / 2
			} else {
				pivot = (PileA.Size() / 2) + 1
			}
			//_________________________________
			// 				  traitement pile B
			if !IsinvSorted(PileB) {
				InvSortB(PileA, PileB, minA, maxA, minB, maxB, low)
				continue outerloop
			}
			//_________________________________
			//				traitement à 2 et 3
			if PileA.Size() <= 3 {
				Less3Sorted(PileA, PileB, minA)
				continue outerloop
			}
			//_______________________________
			//				traitement less 6
			if PileA.Size() <= 6 {
				Less6Sorted(PileA, PileB, minA, maxA, low, i, n, pivot)
				continue outerloop
			}
			// //_______________________________
			// //				  traitement à 10
			// if PileA.Size() <= 10 {
			// 	fmt.Println("defal nank")
			// }
			// //_______________________________
			// //				  traitement à 50
			// if PileA.Size() <= 50 {
			// 	fmt.Println("defal nank")
			// }
			// //_______________________________
			// //				  traitement à 10
			// if PileA.Size() <= 100 {
			// 	fmt.Println("defal nank")
			// }
		}
	}
	if IsSorted(PileA) && len(*PileB) != 0 {
		cmd.Pa(PileB, PileA)
		fmt.Println("Pa")
		models.Count++
		Swapper(PileA, PileB)
	}
}

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
			} else if (n == l && i == 0) && PileB.Size() != 0 {
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
			fmt.Println("low:=", l)
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

// cette fonction permet de trier la pile B
func InvSortB(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, minB int, maxB int, low []int) {
	fmt.Println("pile b my nigger")
	//_________________________
	// 			possible 2 move
	// a: 9 6 5 , b: 1 3 2
	if maxA == 0 && minB == 0 {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	// a: 6 5 9 , b:  2 3 1
	if maxB == 1 && minA == 1 {
		cmd.Ss(PileA, PileB)
		fmt.Println("Ss")
		models.Count++
		return
	}
	if PileB.Size() >= 2 && (maxA == len(*PileA)-1 && minB == len(*PileB)-1) {
		cmd.Rrr(PileA, PileB)
		fmt.Println("Rrr")
		models.Count++
		return
	}
	if PileB.Size() > 2 && (maxA == 0 && minB == 1) {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	if PileB.Size() > 2 && (maxA == 1 && minB == 1) {
		cmd.Rrr(PileA, PileB)
		fmt.Println("Rrr")
		models.Count++
		return
	}
	// a: 6 8 5 , b :1 0 2
	// deux valeurs
	if PileB.Size() == 2 {
		cmd.Sb(PileB)
		fmt.Println("Sb")
		models.Count++
		return
	}
	// 3 valeurs une seul cas possible car deja trier à deux valeurs
	if PileB.Size() == 3 && minB == 0 {
		// -1 2 1
		cmd.Rb(PileB)
		fmt.Println("Rb")
		models.Count++
		return
	}
	if PileB.Size() > 3 {
		// minB à la pos[0]
		if minB == 0 {
			cmd.Rb(PileB)
			fmt.Println("Rb")
			models.Count++
			return
		}
		// maxB  à la pos[1]
		if maxB == 1 {
			cmd.Sb(PileB)
			fmt.Println("Sb")
			models.Count++
			return
		}
	}
}

// func LesstenSorted() {
// 	if minA == 0 {
// 		cmd.Pb(PileA, PileB)
// 		fmt.Println("Pb")
// 		models.Count++
// 		return
// 	}
// 	// if low == 0 pb
// 	if minA == 1 && maxA != 0 {
// 		cmd.Sa(PileA)
// 		fmt.Println("Sa")
// 		models.Count++
// 		return
// 	}
// 	if maxA == 0 && minA != len(PileA)-1 {
// 		cmd.Ra(PileA)
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
// 	if minA >= pivot {
// 		cmd.Rra(PileA)
// 		fmt.Println("Rra")
// 		models.Count++
// 		return
// 	}
// }
