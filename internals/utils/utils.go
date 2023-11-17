package utils

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

// à supprimer je pense (edieng)
func Check_lower(a *models.Stack) {
	taille := a.Size()
	if taille%2 == 0 {
		low := 1000000000000
		for _, ch := range *a {
			if ch.(int) < low {
				low = ch.(int)
			}
		}
		fmt.Println(low)
	}
	if taille%2 == 1 {
		low := 10000000000
		for _, ch := range *a {
			if ch.(int) < low {
				low = ch.(int)
			}
		}
		fmt.Println(low)
	}
	fmt.Println(*a)
}

//	permet dee checker les plus petites valeurs données par l'user
//
// à supprimer je pense (edieng)
func Check_low(a []int) []int {
	taille := len(a)
	sort.Ints(a)
	var low []int
	if taille%2 == 0 {

		for i, ch := range a {
			if i < taille/2 {
				low = append(low, ch)
			}
		}
		fmt.Println(low)
	}
	if taille%2 == 1 {
		for i, ch := range a {
			if i <= (taille/2)-1 {
				low = append(low, ch)
			}
		}
		fmt.Println(low)
	}
	return low
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

// this function has the role of swapping until having all the intergers sorted

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

// permet d'extraire le min du  pile A
func MinA(a *models.Stack) int {
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

// ceci n'est pas une fonction
func Swapping(PileA *models.Stack, PileB *models.Stack, low []int) {
	var P Positions
outerloop:
	for !IsSorted(PileA) {
		for i, n := range *PileA {
			CheckPos(PileA, low, &P)
			min := MinA(PileA)
			for _, l := range low {
				if n == l && i == 0 {
					cmd.Pb(PileA, PileB)
					fmt.Println("Pb")
					fmt.Println(*PileA...)
					models.Count++
					continue outerloop
				}
				if n == l && i == len(*PileA)-1 {
					cmd.Rra(PileA)
					fmt.Println("Rra")
					models.Count++
					continue outerloop
				}
				// _________
				//       2 1
				if PileA.Size() == 2 {
					cmd.Sa(PileA)
					fmt.Println("Sa")
					models.Count++
					continue outerloop
				} else if PileA.Size() == 3 {
					// ____________
					//        2 1 3
					if n == l && i == 1 && (*PileA)[2].(int) > (*PileA)[0].(int) {
						cmd.Sa(PileA)
						fmt.Println("Sa")
						models.Count++
						continue outerloop
					}
					// ____________
					//        3 1 2
					if n == l && i == 1 && (*PileA)[0].(int) > (*PileA)[2].(int) {
						cmd.Ra(PileA)
						fmt.Println("Ra")
						models.Count++
						continue outerloop
					}
					// ____________
					//        2 3 1
					if n == l && i == 2 && (*PileA)[1].(int) > (*PileA)[2].(int) {
						cmd.Rra(PileA)
						fmt.Println("Rra")
						models.Count++
						continue outerloop
					}
					// _____________
					//         3 2 1
					if n == l && i == 2 && (*PileA)[0].(int) > (*PileA)[1].(int) {
						cmd.Ra(PileA)
						fmt.Println("Ra")
						models.Count++
						continue outerloop
					}
					// _____________
					//         1 3 2
					if n == l && i == 0 && (*PileA)[1].(int) > (*PileA)[2].(int) {
						cmd.Rra(PileA)
						fmt.Println("Rra")
						models.Count++
						continue outerloop
					}
				}
			}
			
			if min == 1 && (*PileA)[2].(int) > (*PileA)[0].(int) {
				cmd.Sa(PileA)
				fmt.Println("Sa")
				models.Count++
				continue outerloop
			}
			// ____________
			//        3 1 2
			if min == 1 && (*PileA)[0].(int) > (*PileA)[2].(int) {
				cmd.Ra(PileA)
				fmt.Println("Ra")
				models.Count++
				continue outerloop
			}
			// ____________
			//        2 3 1
			if min == 2 && (*PileA)[1].(int) > (*PileA)[2].(int) {
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				continue outerloop
			}
			// _____________
			//         3 2 1
			if min == 2 && (*PileA)[0].(int) > (*PileA)[1].(int) {
				cmd.Ra(PileA)
				fmt.Println("Ra")
				models.Count++
				continue outerloop
			}
			// _____________
			//         1 3 2
			if min == 0 && (*PileA)[1].(int) > (*PileA)[2].(int) {
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				continue outerloop
			}
			//min est plus proche du bas
			if min == len(*PileA)-3 || min == len(*PileA)-2 {
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				continue outerloop
			}
		}
	}
	//______________________________________________________
	//           si PA asc et PB desc à verifier si PB != nil
	if IsSorted(PileA) {
		if len(*PileB) > 0 {
			fmt.Println("here")
			cmd.Pa(PileB, PileA)
			fmt.Println("Pa")
			models.Count++
		}
	}
}
