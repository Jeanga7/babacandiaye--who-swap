package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

// cette fonction permet de faire le trie sur B
func SortB(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, minB int, maxB int, low []int) {
	//________________________________
	//						cas 2 move
	if maxA == 0 && minB == 0 {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	// a: 6 5 9 , b:  2 3 1
	if maxB == 2 && minA == 2 && PileB.Size() <= 3 {
		cmd.Ss(PileA, PileB)
		fmt.Println("Ss")
		models.Count++
		return
	}
	// minA == dern pos et maxB == dern pos
	if PileB.Size() >= 2 && (maxB == len(*PileA)-1 && minA == len(*PileB)-1) {
		cmd.Rrr(PileA, PileB)
		fmt.Println("Rrr")
		models.Count++
		return
	}
	//si maxA == 1 && minB == 1 ie possibilité de les monter
	if PileB.Size() > 2 && (maxA == 1 && minB == 1) {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	// maxA pos[0] et minB pos[1] possibilité de mettre maxA à len et minB à 0
	if PileB.Size() > 2 && (maxA == 0 && minB == 1) {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	//________________________________
	//						cas simple
	if PileB.Size() == 2 {
		cmd.Sb(PileB)
		fmt.Println("Sb")
		models.Count++
		return
	} else if PileB.Size() == 3 {
		if minB == 0 {
			cmd.Rb(PileB)
			fmt.Println("Rb")
			models.Count++
			return
		} else if (*PileB)[0].(int) < (*PileB)[1].(int) {
			cmd.Sb(PileB)
			fmt.Println("Sb")
			models.Count++
		}
	} else {
		var stackai *models.Stack
		if maxB == 1 && minB != 0 {
			cmd.Sb(PileB)
			fmt.Println("Sb")
			models.Count++
			return
		}
		if minB == 0 {
			cmd.Rb(PileB)
			fmt.Println("Rb")
			models.Count++
			return
		}
		if (*PileB)[0].(int) < (*PileB)[1].(int) {
			cmd.Sb(PileB)
			fmt.Println("Sb")
			models.Count++
			return
		}
		if maxB != 0 && maxB != 1 {
			stackai = CheckAi(PileB, maxB)
			if IsinvSorted(stackai) && maxB != 0 {
				cmd.Rrb(PileB)
				fmt.Println("Rrb")
				models.Count++
				return
			} else {
				cmd.Rb(PileB)
				fmt.Println("Rb")
				models.Count++
				return
			}
		}
		if maxB == 0 {
			cmd.Rb(PileB)
			fmt.Println("Rb")
			models.Count++
			return
		}
	}
}

// cette fonction recupere la valeur se situant à l'indice 0 jusqu'a maxB
func CheckAi(PileB *models.Stack, maxB int) *models.Stack {
	var StackAi models.Stack
	for i := 0; i < maxB; i++ {
		StackAi = append(StackAi, (*PileB)[i])
	}
	return &StackAi
}
