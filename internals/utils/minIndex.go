package utils

import "push/internals/models"

//________________________________________________MinIndex returns the index of the minimum 
//                                                element in the stack.
func MinIndex(slice models.Stack) int {
	if len(slice) == 0 {
		return -1
	}

	minIndex := 0
	minValue := slice[0].(int)
	for i, value := range slice {
		if value.(int) < minValue {
			minIndex = i
			minValue = value.(int)
		}
	}

	return minIndex
}
