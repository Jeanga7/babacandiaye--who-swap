package utils

import (
	"fmt"
	"os"
	"push/internals/models"
)

/*
__________________________________________________________The function "EstTrie" checks

	if a stack is sorted in ascending order.
*/
func EstTrie(pileA, pileB *models.Stack) bool {
	for i := 0; i < len(*pileA)-1; i++ {
		if (*pileA)[i].(int) > (*pileA)[i+1].(int) {
			return false
		}
	}

	// Vérifier si pileB est vide
	if len(*pileB) == 0 {
		fmt.Println("Déjà triée :", *pileA, *pileB)
		fmt.Println("    ///////////////////////////////////////////")
		fmt.Println("   ///                                     ///")
		fmt.Printf("  ///    NOMBRE TOTAL DE MOUVEMENTS: %v    ///\n", Count)
		fmt.Println(" ///                                     ///")
		fmt.Println("///////////////////////////////////////////")
		os.Exit(0)

	}

	// Si pileA est triée et pileB est vide, imprimer le message et arrêter le processus
	return false
}
