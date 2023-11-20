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
			if !IsinvSorted(PileB) {
				SortB(PileA, PileB, minA, maxA, minB, maxB, low)
				continue outerloop
			}
			if PileA.Size() <= 3 {
				Less3Sorted(PileA, PileB, minA)
				continue outerloop
			}
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

