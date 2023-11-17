package utils

import (
	"push/internals/models"
)

// EstPartieTriee vérifie si une partie du stack est triée
func EstPartieTriee(pile *models.Stack, debut, fin int) bool {
	if debut < 0 || fin >= len(*pile) {
		return false
	}

	for i := debut; i < fin; i++ {
		if (*pile)[i].(int) > (*pile)[i+1].(int) {
			return false
		}
	}

	return true
}
