// Package slice contains an implementation of a stack backed by a slice.
package slice

// Stack represents a stack.
type Stack[T any] struct {
	e []T
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.e)
}

// Empty returns whether the stack is empty.
func (s *Stack[T]) Empty() bool {
	return len(s.e) == 0
}

// Push adds an element onto the top of the stack.
func (s *Stack[T]) Push(x T) {
	s.e = append(s.e, x)
}

// Pop removes and returns the element at the top of the stack.
// If the stack is empty - default value for element's type is returned.
func (s *Stack[T]) Pop() T {
	var x T

	if s.Empty() {
		return x
	}

	n := len(s.e)

	x = s.e[n-1]
	s.e = s.e[:n-1]

	return x
}

// Peek returns the element at the top of the stack.
// If the stack is empty - default value for element's type is returned.
func (s *Stack[T]) Peek() T {
	var x T

	if s.Empty() {
		return x
	}

	return s.e[len(s.e)-1]
}
