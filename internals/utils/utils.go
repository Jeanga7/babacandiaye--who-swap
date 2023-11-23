package utils

import (
	"fmt"
	"os"
	"push/internals/algorithm"
	"push/internals/cmd"
	"push/internals/models"
	"runtime"
	"sort"
	"strings"
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

// permet dee checker les plus petites valeurs données par l'user
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

// verifie la validité des arguments
func IsValidArg(arg string) {
	if strings.Contains(arg, " ") {
		for _, ch := range arg {
			if (ch >= '0' || ch <= '9') || ch == 32 {
				continue
			} else {
				fmt.Println("Error")
				return
			}
		}
	} else {
		fmt.Println("\033[91m" + "Error : " + "\033[0m" + "At least 2 numbers are required")
		os.Exit(0)
	}
}

// cette fonction permet de recuperer les piles, de checker les potentiels erreurs et de les tries s'ils sont correctes
func Traitement(A_number models.Stack, B_number models.Stack, en []int) {
	if IsSorted(&A_number) {
		fmt.Println("\033[91m" + "Error" + "\033[0m" + ": Sorted Stack")
		return
	}
	if A_number.Size() > 50 {
		fmt.Println("\033[91m" + "Error" + "\033[0m" + ": The Algorithm doesn't sort the number greater than 50")
		return
	}
	lows := Check_low(en)
	fmt.Print("Stack A : ")
	fmt.Println(A_number...)
	fmt.Println("=======P=U=S=H==S=W=A=P=======")
	algorithm.Swapper(&A_number, &B_number, lows)
	fmt.Println("this is the new positions")
	fmt.Println(A_number...)
	fmt.Println("\033[91m"+"Move :"+"\033[0m", models.Count)
}

// cette fonction permet de checker s'il y'a presence d'une doublure dans la pile
func CheckDuplication(nombre int, pile models.Stack) error {
	for _, n := range pile {
		if n == nombre {
			return fmt.Errorf("error : Le nombre %v est déjà présent dans la pile.", nombre)
		}
	}
	return nil
}
