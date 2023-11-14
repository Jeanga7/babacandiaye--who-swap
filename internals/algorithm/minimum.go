package algorithm

import "fmt"

func FindMin(pile []int) (int, int) {
	// Vérifiez si la pile est vide
	if len(pile) == 0 {
		fmt.Println("Erreur")
	}
	pos := 0
	min := pile[0]
	for i, nombre := range pile {
		if nombre < min {
			min = nombre
			pos = i
		}
	}

	return min, pos
}

// func main() {
// 	pile := []int{3, 1, 4, -1, 5}

// 	minimum, pos := FindMin(pile)
// 	fmt.Println("Minimum dans la pile:", minimum)
// 	fmt.Println("Minimum est à la pos:", pos)
// }
