package main

import (
	"fmt"
	"os"
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
		utils.IsValidArg(arg)
		entier := strings.Split(arg, " ")
		var en []int
		for _, v := range entier {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("\033[91m" + "Error" + "\033[0m" + ": Not a Number")
				os.Exit(0)
			}
			err = utils.CheckDuplication(n, A_number)
			if err != nil {
				fmt.Println("\033[91m" + "Error" + "\033[0m" + ": Duplicate Number")
				os.Exit(0)
			}
			en = append(en, n)
			A_number = append(A_number, n)
		}
		for _, v := range A_number {
			models.S.Push(v)
		}

		utils.Traitement(A_number, B_number, en)

	} else {
		fmt.Println("\033[91m" + "Error" + "\033[0m" + ": Usage : go run . 'number'")
		fmt.Println("       Example : go run . '2 1 3'")
		return
	}
}
