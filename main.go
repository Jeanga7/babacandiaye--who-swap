package main

import (
	"fmt"
	"os"
	"push/internals/cmd"
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

		for _, v := range entier {
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
		fmt.Println("//////////////////////////////////////////////////////")
		fmt.Println("PB ")
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
		fmt.Println("//////////////////////////////////////////////////////")
		fmt.Println("PA")
		cmd.Pb(&A_number, &B_number)
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
		cmd.Sa(&A_number)
		fmt.Println(A_number...)
		fmt.Println("/////////////////////////A_number")
		cmd.Ra(&A_number)
		fmt.Println(A_number...)
		// fmt.Println("//////////////////////")
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
