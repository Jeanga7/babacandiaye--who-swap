package utils

import ("push/internals/models"
"math")

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

func SecondMinIndex(slice models.Stack, excludeIndex int) int {
	if len(slice) == 0 {
		return -1
	}

	secondMinIndex := -1
	secondMinValue := math.MaxInt

	for i, value := range slice {
		if i != excludeIndex && value.(int) < secondMinValue {
			secondMinIndex = i
			secondMinValue = value.(int)
		}
	}

	return secondMinIndex
}

