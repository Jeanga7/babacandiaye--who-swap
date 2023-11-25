package utils

import (
	"fmt"
	"push/internals/models"
)

// Fonction principale pour trier la pile d'entiers entre 10 et 100
func RadixSort(pile *models.Stack) {
	// Trouver le nombre maximal de chiffres pour déterminer le nombre d'itérations
	maxIndex := MaxIndex(*pile)
	maxNum := (*pile)[maxIndex].(int)
	fmt.Println("maxnum", maxNum)
	numDigits := CountDigits(maxNum)

	// Initialisation des seaux
	buckets := make([][]int, 10)

	// Itérations pour chaque chiffre
	for i := 0; i < numDigits; i++ {
		// Placer les éléments dans les seaux en fonction du chiffre actuel
		buckets = DistributeElements(*pile, buckets, i)

		// Reconstituer la pile en combinant les éléments des seaux
		*pile = RebuildList(buckets)

		// Réinitialiser les seaux pour la prochaine itération
		buckets = ResetBuckets(buckets)
	}

	// La pile est maintenant triée
	fmt.Println("Pile triée :", *pile)
}



// Fonction pour compter le nombre de chiffres d'un nombre
func CountDigits(num int) int {
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

// Fonction pour placer les éléments dans les seaux en fonction du chiffre actuel
func DistributeElements(pile models.Stack, buckets [][]int, digitPlace int) [][]int {
	for _, num := range pile {
		digit := (num.(int) / PowerOf10(digitPlace)) % 10
		buckets[digit] = append(buckets[digit], num.(int))
	}
	return buckets
}

// Fonction pour reconstituer la pile en combinant les éléments des seaux
func RebuildList(buckets [][]int) models.Stack {
	var result models.Stack
	for _, bucket := range buckets {
		// Convertir chaque élément de bucket en type interface{}
		var bucketInterface []interface{}
		for _, num := range bucket {
			bucketInterface = append(bucketInterface, num)
		}
		result = append(result, bucketInterface...)
	}
	return result
}
// Fonction pour réinitialiser les seaux pour la prochaine itération
func ResetBuckets(buckets [][]int) [][]int {
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	return buckets
}

// Fonction pour calculer 10 puissance n
func PowerOf10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
