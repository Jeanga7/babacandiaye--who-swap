package utils

import (
	"math"
	"push/internals/models"
)

func SecondMinIndex(slice models.Stack, minIndex int) int {
	if len(slice) == 0 {
		return -1
	}

	secondMinIndex := -1
	secondMinValue := math.MaxInt64

	for i, value := range slice {
		if i == minIndex {
			// Ignorer l'élément avec l'index du plus petit
			continue
		}

		if value.(int) < secondMinValue {
			secondMinIndex = i
			secondMinValue = value.(int)
		}
	}

	return secondMinIndex
}
