package utils

import "push/internals/models"

//_____________________________________________________MaxIndex retourne l'index du plus grand
//                                                     élément dans la pile.
func MaxIndex(slice models.Stack) int {
	if len(slice) == 0 {
		return -1 // Retourne -1 si la pile est vide.
	}

	maxIndex := 0
	maxValue := slice[0].(int)

	for i, value := range slice {
		if value.(int) > maxValue {
			maxIndex = i
			maxValue = value.(int)
		}
	}

	return maxIndex
}

