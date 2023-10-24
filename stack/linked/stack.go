// Package linked contains implementation of a stack backed by linked list.
package linked

import "github.com/denpeshkov/datastructures/list"

// Stack represents a stack.
type Stack[T any] struct {
	l *list.Linked[T]
}

// New returns an initialized stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{list.New[T]()}
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return s.l.Len()
}

// Empty returns whether the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.l.Empty()
}

// Push adds an element onto the top of the stack.
func (s *Stack[T]) Push(x T) {
	s.l.InsertBack(x)
}

// Pop removes and returns the element at the top of the stack.
// If the stack is empty - default value for element's type is returned.
func (s *Stack[T]) Pop() T {
	if s.l.Empty() {
		return *new(T)
	}
	e := s.l.Back()
	s.l.Remove(e)
	return e.Value
}

// Peek returns the element at the top of the stack.
// If the stack is empty - default value for element's type is returned.
func (s *Stack[T]) Peek() T {
	if s.l.Empty() {
		return *new(T)
	}
	return s.l.Back().Value
}
