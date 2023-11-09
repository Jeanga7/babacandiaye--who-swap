package main

import (
	"fmt"
	"os"
	"push/internals/cmd"
	"push/internals/models"
	"push/internals/utils"
	"strconv"
)

func main() {

	if len(os.Args) == 2 {
		arg := os.Args[1]
		var A_number models.Stack
		var B_number models.Stack
		for _, v := range arg {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("Error")
				return
			}
			A_number = append(A_number, n)
		}

		// sort.Ints(A_number)

		for _, v := range A_number {
			models.S.Push(v)
		}
		for i, n := range A_number {
			fmt.Printf("%d\n", n)
			if i == (A_number.Size())-1 {
				fmt.Println("=\t=")
				fmt.Println("a\tb")
			}
		}
		fmt.Println("//////////////////////////////////////////////////////")
		fmt.Println("After popping the top element ")

		cmd.Pb(&A_number, &B_number)
		cmd.Pa(&A_number, &B_number)

		for i, n := range A_number {
			var B_nb string
			if i < len(B_number) {
				B_nb = fmt.Sprintf("%d", B_number[i])
			} else {
				B_nb = ""
			}
			fmt.Printf("%d\t%s\n", n, B_nb)
			if i == (A_number.Size())-1 {
				fmt.Println("=\t=")
				fmt.Println("a\tb")
			}
		}
		fmt.Println(cmd.Sa(A_number))
		// res := Sa(&arg)
		// fmt".Println(res)
	} else {
		utils.Throw("Need a Argument" + utils.Log())
	}
}
