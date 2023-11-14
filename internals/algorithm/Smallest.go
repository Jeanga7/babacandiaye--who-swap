package algorithm

import (
	"fmt"
	"sort"
)

func Smallest(numbers []int) {

	numToExtract := len(numbers) / 2
	if numToExtract < 0 {
		numToExtract = 0
	}
	sort.Ints(numbers)
	smallestNumbers := numbers[:numToExtract]
	fmt.Printf("Suite originale : %v\n", numbers)
	fmt.Printf("Plus petits nombres : %v\n", smallestNumbers)
}

// func main() {
// 	numbers := []int{9, 3, 5, 15, 99, 1}
// 	Smallest(numbers)
// }
