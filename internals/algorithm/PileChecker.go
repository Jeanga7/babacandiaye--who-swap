package algorithm

// Fonction pour vérifier si une pile est ordonnée
func isOk(stack []int) bool {
	length := len(stack)
	// Si la pile est vide ou a un seul élément, elle est considérée comme ordonnée warning×
	if length <= 1 {
		return true
	}

	for i := 1; i < length; i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}

	return true
}

// func main() {
// 	stack := []int{1, 9, 3, 4, 5}

// 	fmt.Printf("La pile est ordonnée : %v\n", isOk(stack))
// }
