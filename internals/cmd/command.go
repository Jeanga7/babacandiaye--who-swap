package cmd

import (
	"push/internals/models"
)

func Pa(b *models.Stack, a *models.Stack) (models.Stack, models.Stack) {
	x := b.Unshift()
	a.Shift(x)
	return *a, *b
}

func Pb(a *models.Stack, b *models.Stack) (models.Stack, models.Stack) {
	x := a.Unshift()
	b.Shift(x)
	return *a, *b
}

func Sa(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[0])
		res = append(res, a[2:]...)
		return res
	}
	return a
}
