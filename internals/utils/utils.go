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

type MinMax struct {
	Min int
	Max int
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

func CheckMinMax(a *models.Stack, m *MinMax) {
	m.Max = -1000
	m.Min = 10000000000
	for _, n := range *a {
		if n.(int) > m.Max {
			m.Max = n.(int)
		}
		if n.(int) < m.Min {
			m.Min = n.(int)
		}
	}
}

// à supprimer je pense
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

// check the lowest numbers given by the user
// à supprimer je pense
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

//Check the position of the lowest numbers in the stack

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
	// fmt.Println(pos.Pos)
}

// this function has the role of swapping until having all the intergers sorted
func Swapper(A_number *models.Stack, B_numbers *models.Stack, low []int) {
	// var B_numbers models.Stack
	var Minmax MinMax
	var P Positions
	// fmt.Println(*A_number...)
	if IsSorted(A_number) && B_numbers.Size() == 0 {
		return
	}
	for i, n := range *A_number {
		var B_nb string
		if i < B_numbers.Size() {
			B_nb = fmt.Sprintf("%d", (*B_numbers)[i])
		} else {
			B_nb = ""
		}
		fmt.Printf("%d\t%s\n", n, B_nb)
		if i == A_number.Size()-1 {
			fmt.Println("=\t=")
			fmt.Println("a\tb")
		}
	}
outerloop:
	for i, n := range *A_number {
		CheckPos(A_number, low, &P)
		// min := MinA(A_number)
		CheckMinMax(A_number, &Minmax)
		for _, l := range low {
			if A_number.Size() == 3 && n == l && i == 1 {
				if (*A_number)[0].(int) > (*A_number)[2].(int) {
					cmd.Ra(A_number)
					fmt.Println("Ra")
					models.Count++
					break outerloop
				}
				if (*A_number)[2].(int) > (*A_number)[0].(int) {
					cmd.Sa(A_number)
					fmt.Println("Sa")
					models.Count++
					break outerloop
				}
			}

			if A_number.Size() == 3 && n == l && i == 0 {
				if (*A_number)[2].(int) < (*A_number)[1].(int) {
					cmd.Rra(A_number)
					fmt.Println("Rra")
					cmd.Sa(A_number)
					models.Count++
					fmt.Println("Sa")
					models.Count++
					break outerloop
				}

			}
			if n == l && i == A_number.Size()-1 {
				cmd.Rra(A_number)
				fmt.Println("Rra")
				models.Count++
				break outerloop
			}
			if (n == l && i == 0) && A_number.Size() >= B_numbers.Size() {
				cmd.Pb(A_number, B_numbers)
				fmt.Println("Pb")
				fmt.Println(*A_number...)
				models.Count++
				if !IsinvSorted(B_numbers) {
					cmd.Sa(B_numbers)
					fmt.Println("Sa")
					models.Count++
					break outerloop
				} else {
					break outerloop
				}
			}
			if n == l && i == 1 {
				cmd.Sa(A_number)
				fmt.Println("Sa")
				fmt.Println("je vais MAMa")
				models.Count++
				break outerloop
			}
			if n == Minmax.Min && i == 1 {
				fmt.Println("je suis min", Minmax.Min)
				cmd.Sa(A_number)
				fmt.Println("Sa")
				fmt.Println("je vais MAMa")
				models.Count++
				break outerloop
			}
		}
		if n == Minmax.Min && i == 1 {
			cmd.Sa(A_number)
			fmt.Println("Sa")
			models.Count++
			break outerloop
		}
		if n == Minmax.Min && i > A_number.Size()/2 {
			cmd.Rra(A_number)
			fmt.Println("Rra")
			models.Count++
			break outerloop
		}
		if !IsSorted(A_number) && n == Minmax.Min && i == 0 {
			cmd.Pb(A_number, B_numbers)
			fmt.Println("Pb")
			models.Count++
			break outerloop
		}
		if IsSorted(A_number) && B_numbers.Size() != 0 {
			cmd.Pa(B_numbers, A_number)
			fmt.Println("Pa")
			models.Count++
			break outerloop
		}
		// if IsinvSorted(A_number) {
		// 	cmd.Rra(A_number)
		// 	fmt.Println("Rra")
		// 	models.Count++
		// 	break outerloop
		// }
		// if IsSorted(B_numbers) && IsinvSorted(A_number) && B_numbers.Size() >= 2 {
		// 	cmd.Ss(A_number, B_numbers)
		// 	fmt.Println("SS")
		// 	models.Count++
		// 	break outerloop
		// }
		// if A_number.Size()%2 == 1 {
		// 	if n == min && i == (A_number.Size()/2)+1 {
		// 		cmd.Rra(A_number)
		// 		fmt.Println("Rra")
		// 		models.Count++
		// 		break outerloop
		// 	}

		// }
		// if IsinvSorted(A_number) {
		// 	for i := 0; i < A_number.Size()-1; i++ {
		// 		cmd.Rra(A_number)
		// 		fmt.Println("Rra")

		// 		models.Count++
		// 		if IsSorted(A_number) {
		// 			return
		// 		}
		// 	}
		// }
	}

	if !IsSorted(A_number) {
		fmt.Println(".....................")
		Swapper(A_number, B_numbers, low)
		// SortedA(A_number)
	} else {
		if IsinvSorted(B_numbers) {
			// loop:
			for i := B_numbers.Size() - 1; i >= 0; i-- {
				cmd.Pa(B_numbers, A_number)
				fmt.Println("Pa")
				models.Count++
				if !IsSorted(A_number) {
					if (*A_number)[1].(int) == Minmax.Min {
						fmt.Println("je suis min", Minmax.Min)
						// fmt.Println("A_number is this")
						cmd.Sa(A_number)
						fmt.Println("Sa")
						fmt.Println("je vais MAMa")
						models.Count++
						for i, n := range *A_number {
							var B_nb string
							if i < B_numbers.Size() {
								B_nb = fmt.Sprintf("%d", (*B_numbers)[i])
							} else {
								B_nb = ""
							}
							fmt.Printf("%d\t%s\n", n, B_nb)
							if i == A_number.Size()-1 {
								fmt.Println("=\t=")
								fmt.Println("a\tb")
							}
						}

					}
				}
			}
			// if IsSorted(A_number) {
			// 	return
			// } else {
			// 	// SwapA(A_number, B_numbers)

			// }

		} else {
			fmt.Println("je fais le SwapB")
			SwapB(A_number, B_numbers)
		}
	}

}

func SwapB(A_number *models.Stack, B_numbers *models.Stack) {
	var Minmax MinMax
	CheckMinMax(B_numbers, &Minmax)
outerloop:
	for i, n := range *B_numbers {
		if !IsinvSorted(B_numbers) {
			if n == Minmax.Max && i == 0 {
				cmd.Pa(B_numbers, A_number)
				fmt.Println("Pa")
				models.Count++
				break outerloop
			}
			if IsinvSorted(B_numbers) && n == Minmax.Max && i == 0 {
				cmd.Pa(B_numbers, A_number)
				fmt.Println("Pa")
				models.Count++
				break outerloop
			}
			if n == Minmax.Max && i == 1 {
				cmd.Sb(B_numbers)
				fmt.Println("Sb")
				models.Count++
				break outerloop
			}
			if B_numbers.Size()%2 == 0 && n == Minmax.Max && i != 0 {
				if n == Minmax.Max && i > B_numbers.Size()/2 {
					cmd.Rrb(B_numbers)
					fmt.Println("Rrb")
					models.Count++
					break outerloop
				}
				if n == Minmax.Max && i <= B_numbers.Size()/2 {
					cmd.Rb(B_numbers)
					fmt.Println("Rb")
					models.Count++
					break outerloop
				}

			}
			if B_numbers.Size()%2 == 1 && n == Minmax.Max && i != 0 {
				if n == Minmax.Max && i > B_numbers.Size()/2 {
					cmd.Rrb(B_numbers)
					fmt.Println("Rrb")
					models.Count++
					break outerloop
				}
				if n == Minmax.Max && i <= B_numbers.Size()/2 {
					cmd.Rb(B_numbers)
					fmt.Println("Rb")
					models.Count++
					break outerloop
				}
			}
		} else {
			if n == Minmax.Max && i == 0 {
				cmd.Pa(B_numbers, A_number)
				fmt.Println("Pa")
				models.Count++
				break outerloop
			}
		}
	}
	if !IsinvSorted(B_numbers) {
		SwapB(A_number, B_numbers)
	}

	// if !IsSorted(A_number) {

	// }
}

func SwapA(A_number *models.Stack, B_numbers *models.Stack) {
	var Minmax MinMax
	CheckMinMax(B_numbers, &Minmax)
outerloop:
	for i, n := range *A_number {
		// if !IsinvSorted(A_number) && n == Minmax.Max && i == 0 {
		// 	cmd.Pa(A_number, B_numbers)
		// 	fmt.Println("Pa")
		// 	models.Count++
		// 	break outerloop
		// }
		if n == Minmax.Min && i == 1 {
			cmd.Sa(A_number)
			fmt.Println("Sa")
			models.Count++
			break outerloop
		}
		if n == Minmax.Min && i > A_number.Size()/2 {
			cmd.Rra(A_number)
			fmt.Println("Rra")
			models.Count++
			break outerloop
		}
		// if n == Minmax.Min && i != A_number.Size()-1 {
		// 	cmd.Rrb(A_number)
		// 	fmt.Println("Rra")
		// 	models.Count++
		// 	break outerloop
		// }
	}
	if !IsinvSorted(B_numbers) {
		SwapB(A_number, B_numbers)
	}
	// if !IsSorted(A_number) {

	// }
}

func IsSorted(stack *models.Stack) bool {
	for i := 0; i < stack.Size()-1; i++ {
		if (*stack)[i].(int) > (*stack)[i+1].(int) {
			return false
		}
	}
	return true
}

func SortedA(stack *models.Stack) {

	if (*stack)[0].(int) > (*stack)[1].(int) {
		cmd.Sa(stack)
	}

}

func IsinvSorted(stack *models.Stack) bool {
	for i := 0; i < stack.Size()-1; i++ {
		if (*stack)[i].(int) < (*stack)[i+1].(int) {
			return false
		}
	}
	return true
}

func MinA(a *models.Stack) int {
	min := 1000000000000
	for _, x := range *a {
		if x.(int) < min {
			min = x.(int)
		}
	}
	return min
}
