package main

import (
	"fmt"
	"os"
	"push/internals/cmd/pushswap"
	"push/internals/models"
	"push/internals/utils"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var A_number *models.Stack = &models.Stack{}
		var B_number *models.Stack = &models.Stack{}
		for _, ch := range arg {
			if strings.Contains(arg, " ") {
				if (ch >= '0' && ch <= '9') || ch == 32 {
					continue
				} else {
					fmt.Println("not a valid argument")
					return
				}
			} else if  !strings.Contains(arg, " ") && (ch >= '0' && ch <= '9') {
				continue
			} else {
				fmt.Println("not a valid argument")
				return
			}
		}
		entier := strings.Split(arg, " ")
		for _, v := range entier {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("Error")
				return
			}
			*A_number = append(*A_number, n)
		}
		if len(*A_number) <= 3 {
			pushswap.TrierTroisNombres(A_number, B_number)
		}
	} else {
		utils.Throw("Need an Argument" + utils.Log())

	}
}
