package algorithm

import (
	"fmt"
)

func FindMax(pile []int) (int,int) {
	// Vérifiez si la pile est vide jusqu'a l'arriver de checker
	if len(pile) == 0 {
		fmt.Println("La pile est vide.")
	}
	pos := 0
	max := pile[0]
	for i, nombre := range pile {
		if nombre > max {
			max = nombre
			pos = i
		}
	}

	return max,pos
}

// func main() {
// 	pile := []int{99, 7, 1, 9, 5, 2, 8, 4, 6}
// 	maximum,pos := FindMax(pile)
// 	fmt.Println("Le maximum dans la pile est:", maximum)
// 	fmt.Println("Le maximum est à la pos:", pos)
// }
