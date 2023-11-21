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

func Sa(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[1]}, (*a)[0])
		res = append(res, (*a)[2:]...)
		*a = res
	}
}
func Sb(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[1]}, (*a)[0])
		res = append(res, (*a)[2:]...)
		*a = res
	}
}
func Ra(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[1]}, (*a)[2:]...)
		res = append(res, (*a)[0])
		*a = res
	}
}
func Rb(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[1]}, (*a)[2:]...)
		res = append(res, (*a)[0])
		*a = res
	}
}

func Rr(a *models.Stack, b *models.Stack) {
	Ra(a)
	Rb(b)
}

func Ss(a *models.Stack, b *models.Stack) {
	Sa(a)
	Sb(b)
}

func Rra(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[len(*a)-1]}, (*a)[:len(*a)-1]...)
		// res = append(res, a[0])
		*a = res
	}
}
func Rrb(a *models.Stack) {
	if len(*a) > 1 {
		res := append([]interface{}{(*a)[len(*a)-1]}, (*a)[:len(*a)-1]...)
		// res = append(res, a[0])
		*a = res
	}
}
func Rrr(a *models.Stack, b *models.Stack) {
	Rra(a)
	Rrb(b)
}
