package models

type Stack []interface {
}

// The `Push` function is used to add an item to the top of the stack.
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// The `Shift` function is used to insert an item at the beginning of the stack.
func (s *Stack) Shift(item interface{}) {
	newStack := make(Stack, len(*s)+1)
	newStack[0] = item
	copy(newStack[1:], *s)
	*s = newStack
}

// remove the top element of the stack
func (s *Stack) Pop() (item interface{}) {
	item, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

func (s *Stack) Unshift() (item interface{}) {
	item, *s = (*s)[0], (*s)[1:]
	return
}

// verify if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// return the top element of the stack otherwise return nil
func (s *Stack) Peek() (item interface{}) {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

// return the length of the stack
func (s *Stack) Size() int {
	return len(*s)
}
