package trier_grande_pile

import (
	"fmt"
	"push/internals/cmd"
	psPetit "push/internals/cmd/pushswap/trier_petit_pile"
	"push/internals/models"
	"push/internals/utils"
)

func TrierDixNombres(pileA, pileB *models.Stack) {
	if utils.EstTrie(pileA, pileB) {
		fmt.Println("Déjà triés 6:", *pileA, *pileB)
		return
	}
	fmt.Println("requiredA", RequireMovesPileA(pileA))
	// fmt.Println("requiredB", RequireMovesPileB(pileA,pileB))

	for pileA.Size() > 5 {
		utils.RemonterPlusPetit(pileA, pileB)
	}

	if pileA.Size() == 5 {
		fmt.Println("debut trie a 5")
		psPetit.TrierCinqNombres(pileA, pileB)
		fmt.Println("fin trie a 5")
	}

	for len(*pileB) > 0 {
		if (*pileA)[0].(int) > (*pileA)[1].(int) {
			cmd.Sa(pileA)
			utils.CounterMoves()
			fmt.Println("Apres sa 6:", *pileA, *pileB)
		}
		cmd.Pa(pileB, pileA)
		utils.CounterMoves()
		fmt.Println("Après pa 6:", *pileA, *pileB)
		psPetit.TrierTroisNombres(pileA, pileB)
	}

	if !utils.EstTrie(pileA, pileB) && pileB.Size() == 0 {
		cmd.Sa(pileA)
		utils.CounterMoves()
		fmt.Println("Apres sa 6:", *pileA, *pileB)
	}
}

//=======================================================================================
func RequireMovesPileA(pileA *models.Stack) []int {
	taillePile := pileA.Size()
	movesNecessaires := make([]int, taillePile)

	for _, nombre := range *pileA {
		nombreIndex := TrouverIndexDansPile(pileA, nombre)
		mouvementsRA := nombreIndex
		mouvementsRRA := taillePile - nombreIndex

		if mouvementsRA < mouvementsRRA {
			movesNecessaires[nombreIndex] = mouvementsRA
		} else {
			movesNecessaires[nombreIndex] = mouvementsRRA
		}
	}

	return movesNecessaires
}

func TrouverIndexDansPile(pileA *models.Stack, nombre interface{}) int {
	for i, num := range *pileA {
		if num == nombre {
			return i
		}
	}
	return -1
}
//=====================================================================
