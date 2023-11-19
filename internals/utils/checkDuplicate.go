package utils

import "push/internals/models"

func CheckDuplicate(pileA *models.Stack, nouveauNombre int) bool {
	for _, nombre := range *pileA {
		if nombre == nouveauNombre {
			return true // Le nombre est déjà présent
		}
	}
	return false // Le nombre n'est pas présent
}
