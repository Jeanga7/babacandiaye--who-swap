package cmd

import (
	"push/internals/models"
)

type Test struct {
	Input  models.Stack
	Output models.Stack
	Want   models.Stack
}

// The TestPa function tests the Pa function by comparing its output with an expected output and
// printing "Pass" if they match.
// func TestPa(t *testing.T) {
// 	sa := []int{2, 1, 3}
// 	expect := []int{1, 2, 3}
// 	if result := Pa(sa); !reflect.DeepEqual(result, expect) {
// 		t.Errorf("Fail expected=%v, got=%v", expect, result)
// 	} else {
// 		fmt.Println("Pass")
// 	}
// }
