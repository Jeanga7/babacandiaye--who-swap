package main

import (
	"fmt"
	"os"
	"push/internals/algorithm"
	"push/internals/models"
	"push/internals/utils"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var A_number models.Stack
		var B_number models.Stack
		IsValidArg(arg)
		entier := strings.Split(arg, " ")
		var en []int
		for _, v := range entier {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("\033[91m"+"Error"+"\033[0m"+": Not a Number")
				return
			}
			err = CheckDuplication(n, A_number)
			if err != nil {
				fmt.Println("\033[91m"+"Error"+"\033[0m"+": Duplicate Number")
				return
			}
			en = append(en, n)
			A_number = append(A_number, n)
		}
		for _, v := range A_number {
			models.S.Push(v)
		}

		Traitement(A_number, B_number, en)

	} else {
		utils.Throw("Need a Argument" + utils.Log())
	}			
}

// __________________________________________________
// 					verifie la validité des arguments
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
		fmt.Println("\033[91m"+"Error"+"\033[0m"+"not a valid argument")
		return
	}
}

func Traitement(A_number models.Stack,B_number models.Stack,en []int){
	var P utils.Positions
	if	utils.IsSorted(&A_number) {
		fmt.Println( "\033[91m"+"Error"+"\033[0m"+": Sorted Stack")
		return
	}
	// CheckDuplication(n, pileA)
	// low := utils.Check_low(en)
	fmt.Println(A_number...)
	// utils.CheckPos(&A_number, low, &P)
	fmt.Println(P)
	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	utils.Check_lower(&A_number)
	// algorithm.Swapping(&A_number, &B_number, low)
	// utils.Swapping(&A_number, &B_number, low)
	algorithm.Swapper(&A_number, &B_number)
	fmt.Println("this is the new positions")
	fmt.Println(A_number...)
	fmt.Println("move :",models.Count)
}

func CheckDuplication(nombre int, pile models.Stack) error {
	for _, n := range pile {
		if n == nombre {
			return fmt.Errorf("error : Le nombre %v est déjà présent dans la pile.", nombre)
		}
	}
	return nil
}
