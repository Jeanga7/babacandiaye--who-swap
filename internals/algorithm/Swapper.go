package algorithm

import (
	"fmt"
	"os"
	"push/internals/cmd"
	"push/internals/models"
	"runtime"
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

// ceci n'est pas une fonction
func Swapping(PileA *models.Stack, PileB *models.Stack, low []int) {
	var P Positions
outerloop:
	for !IsSorted(PileA) {
		for i, n := range *PileA {
			CheckPos(PileA, low, &P)
			minA := MinPile(PileA)
			maxA := MaxPile(PileA)
			minB := MinPile(PileB)
			maxB := MaxPile(PileB)
			//_____________________________________
			// 			traitement à 2 et 3 (pileA)
			if PileA.Size() <= 3 {
				if PileA.Size() == 2 {
					cmd.Sa(PileA)
					fmt.Println("Sa")
					models.Count++
					continue outerloop
				} else if PileA.Size() == 3 {
					// ____________
					//        2 1 3
					if minA == 1 && (*PileA)[2].(int) > (*PileA)[0].(int) {
						cmd.Sa(PileA)
						fmt.Println("Sa")
						models.Count++
						continue outerloop
					}
					// ____________
					//        3 1 2
					if minA == 1 && (*PileA)[0].(int) > (*PileA)[2].(int) {
						cmd.Ra(PileA)
						fmt.Println("Ra")
						models.Count++
						continue outerloop
					}
					// ____________
					// 		3 2 1
					if (minA == 2 && (*PileA)[0].(int) > (*PileA)[1].(int)) && IsinvSorted(PileA) {
						cmd.Ra(PileA)
						fmt.Println("Ra")
						models.Count++
						continue outerloop
					}
					// _______________________
					//        2 3 1  ou  1 3 2
					if (minA == 2 && (*PileA)[1].(int) > (*PileA)[2].(int)) || (minA == 0 && (*PileA)[1].(int) > (*PileA)[2].(int)) {
						fmt.Println(*PileA...)
						cmd.Rra(PileA)
						fmt.Println("Rra")
						models.Count++
						continue outerloop
					}
				}
			}
			//____________________________
			// 			traitement pile B
			if !IsinvSorted(PileB) {
				//_________________________
				// 			possible 2 move
				// a: 9 6 5 , b: 1 3 2
				if maxA == 0 && minB == 0 {
					cmd.Rr(PileA, PileB)
					fmt.Println("Rr")
					models.Count++
					continue outerloop
				}
				// a: 6 5 9 , b:  2 3 1
				if maxB == 1 && minA == 1 {
					cmd.Ss(PileA, PileB)
					fmt.Println("Ss")
					models.Count++
					continue outerloop
				}
				// a: 6 8 5 , b :1 0 2
				// deux valeurs
				if PileB.Size() == 2 {
					cmd.Sb(PileB)
					fmt.Println("Sb")
					models.Count++
					continue outerloop
				}
				// 3 valeurs une seul cas possible car dejas trier à deux valeurs
				if PileB.Size() == 3 && minB == 0 {
					// -1 2 1
					cmd.Rb(PileB)
					fmt.Println("Rb")
					models.Count++
					continue outerloop
				}
				if PileB.Size() > 3 {
					// minB à la pos[0]
					if minB == 0 {
						cmd.Rb(PileB)
						fmt.Println("Rb")
						models.Count++
						continue outerloop
					}
					// maxB  à la pos[1]
					if maxB == 1 {
						cmd.Sb(PileB)
						fmt.Println("Sb")
						models.Count++
						continue outerloop
					}
				}
			}
			//____________________________________________________________
			// 			algorithm de traitement pileA max 8 pour le moment
			
			// minA à la deuxiem pos
			if minA == 1 && PileA.Size() != 3 {
				cmd.Sa(PileA)
				fmt.Println("Sa")
				models.Count++
				continue outerloop
			}
			// minA à la derniere pos et gere aussi si une pile est descendant
			if minA == len(*PileA)-1 {
				cmd.Rra(PileA)
				fmt.Println("ooops")
				fmt.Println("Rra")
				models.Count++
				continue outerloop
			}
			// minA à avant dernier position
			if minA == len(*PileA)-2 {
				cmd.Rra(PileA)
				fmt.Println("Rra")
				models.Count++
				continue outerloop
			}
			// maxA à la pos 0
			if maxA == 0 && !IsinvSorted(PileA) {
				cmd.Ra(PileA)
				fmt.Println("Ra")
				models.Count++
				continue outerloop
			}
			//min se trouve en bas du pivot
			for i := PileA.Size(); i > PileA.Size()/2; i-- {
				if i == minA {
					cmd.Rra(PileA)
					fmt.Println("Rra")
					models.Count++
					continue outerloop
				}
			}
			// min se trouve en haut du pivot
			for i := 1; i <= PileA.Size()/2; i++ {
				if i == minA {
					cmd.Ra(PileA)
					fmt.Println("Ra")
					models.Count++
					continue outerloop
				}
			}
			//__________________________________________________
			// 			traitement à plusieurs limite (10) pileA
			for _, l := range low {
				// j'enleve le pilsize > 3 au cas ou
				if n == l && i == 0 {
					cmd.Pb(PileA, PileB)
					fmt.Println("Pb")
					fmt.Println(*PileA...)
					fmt.Println(*PileB...)
					models.Count++
					continue outerloop
				}
				if n == l && i == len(*PileA)-1 {
					cmd.Rra(PileA)
					fmt.Println("Rra")
					models.Count++
					continue outerloop
				}
			}
			// maxA à avant dernier position
			if maxA == len(*PileA)-2 && minA == 0 {
				cmd.Pb(PileA, PileB)
				fmt.Println("Pb")
				fmt.Println(*PileA...)
				fmt.Println(*PileB...)
				models.Count++
				continue outerloop
			}
			// si max est proche bas et min en haut ie l'interieur n'est pas trier
			for i := PileA.Size(); i > PileA.Size()/2; i-- {
				if i == maxA && minA == 0 {
					cmd.Pb(PileA, PileB)
					fmt.Println("Pb")
					fmt.Println(*PileA...)
					fmt.Println(*PileB...)
					models.Count++
					continue outerloop
				}
			}
		}
	}
	if IsSorted(PileA) && len(*PileB) != 0 {
		cmd.Pa(PileB, PileA)
		fmt.Println(*PileB...)
		fmt.Println("Pa")
		models.Count++
		Swapping(PileA, PileB, low)
	}
}
