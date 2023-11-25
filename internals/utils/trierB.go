package utils

import (
	"fmt"
	"push/internals/cmd"
	"push/internals/models"
)

func TrierB(pileA, pileB *models.Stack) {
	maxIndexA := MaxIndex(*pileA)
	minIndexA := MinIndex(*pileA)
	maxIndexB := MaxIndex(*pileB)
	minIndexB := MinIndex(*pileB)
	
	if pileA.Size() < 2  || pileB.Size() < 2  {
		fmt.Println("Pas besoin de trier triB", *pileA, *pileB)
		return
	}
	if pileB.Size() == 2 {
		if minIndexB == 0 && maxIndexA == 1 && minIndexA == 0 {
			cmd.Ss(pileA, pileB)
			CounterMoves()
			fmt.Println("Après ss triB:", *pileA, *pileB)
		} else if minIndexB == 0 && maxIndexA == 0 && minIndexA == pileB.Size()-1 {
			cmd.Ss(pileA, pileB)
			CounterMoves()
			fmt.Println("Après ss triB:", *pileA, *pileB)
		} else if minIndexB == 0 && minIndexA == pileA.Size()-1 {
			cmd.Rrr(pileA, pileB)
			CounterMoves()
			fmt.Println("Après rra triB:", *pileA, *pileB)
		} else if minIndexB == 0 {
			cmd.Sb(pileB)
			CounterMoves()
			fmt.Println("Après sb triB:", *pileA, *pileB)
		} else {
			EstTrie(pileA, pileB)
		}
		return
	} else {
		if maxIndexB == pileB.Size()-1 && minIndexB == 1 {
			if minIndexA == pileA.Size()-1 {
				cmd.Rrr(pileA, pileB)
				CounterMoves()
				fmt.Println("Après rrr trieB :", *pileA, *pileB)
			} else {
				cmd.Rrb(pileB)
				CounterMoves()
				fmt.Println("Après rrb trieB :", *pileA, *pileB)
			}
		} else if maxIndexB == 1 && minIndexA == pileB.Size()-1 {
			if minIndexA == 1 || (maxIndexA == 0 && minIndexA == pileA.Size()-1) {
				cmd.Ss(pileA, pileB)
				CounterMoves()
				fmt.Println("Après ss trieB :", *pileA, *pileB)
			} else {
				cmd.Sb(pileB)
				CounterMoves()
				fmt.Println("Après sb trieB ici:", *pileA, *pileB)
			}
		} else if maxIndexB == 0 && minIndexB == 1 {
			if minIndexA == 1 {
				cmd.Sb(pileB)
				CounterMoves()
				fmt.Println("Après sb trieB :", *pileA, *pileB)
				cmd.Rr(pileA, pileB)
				CounterMoves()
				fmt.Println("Après rr trieB :", *pileA, *pileB)
			} else {
				cmd.Sb(pileB)
				CounterMoves()
				fmt.Println("Après sb trieB :", *pileA, *pileB)
				cmd.Rb(pileB)
				CounterMoves()
				fmt.Println("Après rb trieB :", *pileA, *pileB)
			}
		} else if maxIndexB == pileB.Size()-1 && minIndexB == 0 {
			if minIndexA == 1 {
				cmd.Ss(pileA, pileB)
				CounterMoves()
				fmt.Println("Après ss trieB :", *pileA, *pileB)
				cmd.Rrb(pileA)
				CounterMoves()
				fmt.Println("Après rrb trieB :", *pileA, *pileB)
			} else if minIndexA == pileA.Size()-1 {
				cmd.Sb(pileB)
				CounterMoves()
				fmt.Println("Après sb trieB :", *pileA, *pileB)
				cmd.Rrr(pileA, pileB)
				CounterMoves()
				fmt.Println("Après rrr trieB :", *pileA, *pileB)
			} else {
				cmd.Sb(pileB)
				CounterMoves()
				fmt.Println("Après sb trieB :", *pileA, *pileB)
				cmd.Rrb(pileB)
				CounterMoves()
				fmt.Println("Après rrb trieB :", *pileA, *pileB)
			}
		} else if maxIndexB == 1 && minIndexB == 0 {
			if minIndexA == 0 {
				cmd.Rr(pileA, pileB)
				CounterMoves()
				fmt.Println("Après rr trieB :", *pileA, *pileB)
			} else {
				cmd.Rb(pileB)
				CounterMoves()
				fmt.Println("Après Rb trieB :", *pileA, *pileB)
			}
		}
	}
	EstTrie(pileA, pileB)
	
}