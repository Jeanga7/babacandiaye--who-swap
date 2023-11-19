package utils

type Stack struct {
	// La définition de votre structure de pile
	data []int
}

// Push ajoute une valeur à la pile
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// GetAllValues renvoie un tableau contenant toutes les valeurs de la pile
func (s *Stack) GetAllValues() []int {
	return s.data
}
