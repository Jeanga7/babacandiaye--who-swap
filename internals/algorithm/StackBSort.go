package algorithm

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

// cette fonction permet de faire le trie sur B 
func SortB(PileA *models.Stack, PileB *models.Stack, minA int, maxA int, minB int, maxB int, low []int) {
	var pivot int
	if PileA.Size()%2 == 0 {
		pivot = PileA.Size() / 2
	} else {
		pivot = (PileA.Size() / 2) + 1
	}
	//________________________________
	//						cas 2 move
	// a: 9 6 5 , b: 1 3 2
	if maxA == 0 && minB == 0 {
		cmd.Rr(PileA, PileB)
		fmt.Println("Rr")
		models.Count++
		return
	}
	// a: 6 5 9 , b:  2 3 1
	if maxB == 1 && minA == 1 {
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
		cmd.Rrr(PileA, PileB)
		fmt.Println("Rrr")
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
	if PileB.Size() == 2 || maxB == 1 {
		cmd.Sb(PileB)
		fmt.Println("Sb")
		models.Count++
		return
	}
	//_________________________________
	//			cas nécéssitant un trie
	if PileB.Size() == 3 && minB == 0 {
		// -1 2 1
		cmd.Rb(PileB)
		fmt.Println("Rb")
		models.Count++
		return
	}
	// maxB proche du bas
	if maxB >= pivot {
		cmd.Rrb(PileB)
		fmt.Println("Rrb")
		models.Count++
		return
	}
	// maxB proche du haut
	if maxB < pivot {
		cmd.Rb(PileB)
		fmt.Println("Rb")
		models.Count++
		return
	}
}
