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
		if strings.Contains(arg, " ") {
			for _, ch := range arg {
				if (ch >= '0' || ch <= '9') || ch == 32 {
					continue
				} else {
					fmt.Println("not a valid argument")
					return
				}
			}
		} else {
			fmt.Println("not a valid argument")
			return
		}

		entier := strings.Split(arg, " ")
		var en []int
		for _, v := range entier {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("Error")
				return
			}
			en = append(en, n)
			A_number = append(A_number, n)
		}
		// sort.Ints(A_number)
		for _, v := range A_number {
			models.S.Push(v)
		}

		var P utils.Positions

		low := utils.Check_low(en)
		// fmt.Println(low)
		// fmt.Println("*******************************************")
		fmt.Println(A_number...)
		utils.CheckPos(&A_number, low, &P)
		fmt.Println(P)
		fmt.Println("///////////////////////////////////////////////////////////////////////////")
		utils.Check_lower(&A_number)

		utils.Swapper(&A_number, &B_number, low)
		// utils.CheckPos(&A_number, low, &P)
		fmt.Println("this is the new positions")
		fmt.Println(A_number...)
		fmt.Println(models.Count)
		// fmt.Println(P)
		// fmt.Println("^^^^^^^^^^^^^^^^^^")
		// utils.Swapper(&A_number, &B_number, low)
		// utils.Swapper(&A_number, &B_number, low)
		// fmt.Println(A_number...)
		// fmt.Println("After SS")
		// fmt.Println(cmd.Ss(cmd.Ra(cmd.Sa(A_number)), cmd.Ra(cmd.Sa(B_number))))
		// fmt.Println("////////////////////////////")
		// a, b := cmd.Ss(cmd.Ra(cmd.Sa(A_number)), cmd.Ra(cmd.Sa(B_number)))
		// fmt.Println("//////////////////////////////:")
		// fmt.Println("After RRR")
		// fmt.Println(cmd.Rrr(a, b))
		// fmt.Println(cmd.Rra(a)...)
		// for i, n := range a {
		// 	var B_nb string
		// 	if i < len(b) {
		// 		B_nb = fmt.Sprintf("%d", b[i])
		// 	} else {
		// 		B_nb = ""
		// 	}
		// 	fmt.Printf("%d\t%s\n", n, B_nb)
		// 	if i == len(a)-1 {
		// 		fmt.Println("=\t=")
		// 		fmt.Println("a\tb")
		// 	}
		// }
	} else {
		utils.Throw("Need a Argument" + utils.Log())
	}
}
