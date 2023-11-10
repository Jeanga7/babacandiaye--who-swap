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

func Sb(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[0])
		res = append(res, a[2:]...)
		return res
	}
	return a
}

//////////////////////

func Ra(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[2:]...)
		res = append(res, a[0])
		return res
	}
	return a
}

func Rb(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[2:]...)
		res = append(res, a[0])
		return res
	}
	return a
}

func Ss(a []interface{}, b []interface{}) ([]interface{}, []interface{}) {
	return Sa(a), Sb(b)
}

func Rr(a []interface{}, b []interface{}) ([]interface{}, []interface{}) {
	return Ra(a), Rb(b)
}

func Rra(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[len(a)-1]}, a[:len(a)-1]...)
		// res = append(res, a[0])
		return res
	}
	return a
}

func Rrb(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[len(a)-1]}, a[:len(a)-1]...)
		// res = append(res, a[0])
		return res
	}
	return a
}

func Rrr(a []interface{}, b []interface{}) ([]interface{}, []interface{}) {
	return Rra(a), Rrb(b)
}
