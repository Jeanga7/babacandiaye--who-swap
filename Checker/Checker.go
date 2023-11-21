package main

import (
	"bufio"
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
		var func_todo []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			func_todo = append(func_todo, scanner.Text())
		}
		// fmt.Println(func_todo)
		fmt.Println(Apllication(func_todo, &A_number, &B_number))
	} else {
		return
	}

}

func Apllication(instructions []string, A *models.Stack, B *models.Stack) string {
	for _, v := range instructions {
		if v == "rra" {
			cmd.Rra(A)
		}
		if v == "ra" {
			cmd.Ra(A)
		}
		if v == "sa" {
			cmd.Sa(A)
		}
		if v == "pa" {
			cmd.Pa(B, A)
		}
		if v == "rrb" {
			cmd.Rrb(B)
		}
		if v == "rb" {
			cmd.Rb(B)
		}
		if v == "sb" {
			cmd.Sb(B)
		}
		if v == "pb" {
			cmd.Pb(A, B)
		}
		if v == "ss" {
			cmd.Ss(A, B)
		}
		if v == "rr" {
			cmd.Rr(A, B)
		}
		if v == "rrr" {
			cmd.Rrr(A, B)
		}
	}
	if utils.IsSorted(A) && B.Size() == 0 {
		return "OK"
	} else {
		return "KO"
	}
}
