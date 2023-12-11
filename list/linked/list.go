// Package linked contains an implementation of a doubly-linked circular list.
package linked

import (
	"fmt"
)

// Element represents an element of the list.
type Element[T any] struct {
	Value      T
	l          *Linked[T]
	prev, next *Element[T]
}

// Next returns the next element or nil.
func (e *Element[T]) Next() *Element[T] {
	if e.l == nil || e.next == &e.l.root {
		return nil
	}

	return e.next
}

// Prev returns the previous element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if e.l == nil || e.prev == &e.l.root {
		return nil
	}
	return e.prev
}

// Iterator is an iterator for a linked list.
type Iterator[T any] struct {
	e *Element[T]
}

// Next returns the next element or nil.
func (i *Iterator[T]) Next() *Element[T] {
	return i.e.Next()
}

// Prev returns the previous element or nil.
func (i *Iterator[T]) Prev() *Element[T] {
	return i.e.Prev()
}

// Linked represents a doubly-linked circular list.
type Linked[T any] struct {
	// Sentinel node
	// Head is root.next, tail is root.prev
	root Element[T]
	len  int
}

// New returns an initialized list.
func New[T any]() *Linked[T] {
	l := new(Linked[T])
	l.len = 0
	l.root.prev = &l.root
	l.root.next = &l.root
	return l
}

// Front returns the first element of the list or nil if the list is empty.
func (l *Linked[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of the list or nil if the list is empty.
func (l *Linked[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insertAfter inserts uninitialized element e after element p and returns e.
func (l *Linked[T]) insertAfter(e, p *Element[T]) *Element[T] {
	e.l = l
	e.prev = p
	e.next = p.next
	e.prev.next = e
	e.next.prev = e

	l.len++

	return e
}

// InsertFront inserts a new element e with value v at the front of the list and returns e.
func (l *Linked[T]) InsertFront(v T) *Element[T] {
	return l.insertAfter(&Element[T]{Value: v}, &l.root)
}

// InsertBack inserts a new element e with value v at the back of the list and returns e.
func (l *Linked[T]) InsertBack(v T) *Element[T] {
	return l.insertAfter(&Element[T]{Value: v}, l.root.prev)
}

// InsertBefore inserts a new element with value v immediately before p and returns e.
// If p is not an element of the list l, nil is returned.
func (l *Linked[T]) InsertBefore(v T, p *Element[T]) *Element[T] {
	if p.l != l {
		return nil
	}
	return l.insertAfter(&Element[T]{Value: v}, p.prev)
}

// InsertAfter inserts a new element with value v immediately after p and returns e.
// If p is not an element of the list l, nil is returned.
func (l *Linked[T]) InsertAfter(v T, p *Element[T]) *Element[T] {
	if p.l != l {
		return nil
	}
	return l.insertAfter(&Element[T]{Value: v}, p)
}

// Remove removes element e from list l if e is an element of list l.
func (l *Linked[T]) Remove(e *Element[T]) {
	if e.l != l {
		return
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = nil // avoid loitering
	e.next = nil // avoid loitering
	e.l = nil    // avoid loitering

	l.len--
}

// At returns an element at index ind.
func (l *Linked[T]) At(ind int) (*Element[T], error) {
	if ind >= l.len {
		return nil, fmt.Errorf("index ind=%v out of bounds: [%v, %v]", ind, 0, l.len-1)
	}

	e := l.root.next
	for i := 0; i < ind; i++ {
		e = e.next
	}
	return e, nil
}

// Len returns the number of elements in the list.
func (l *Linked[T]) Len() int {
	return l.len
}

// Empty returns whether the list is empty.
func (l *Linked[T]) Empty() bool {
	return l.len == 0
}

// Iter returns an iterator pointing at the first element.
func (l *Linked[T]) Iter() Iterator[T] {
	return Iterator[T]{l.root.next}
}

// Find returns the first element with the value v, or nil if not present.
// The second parameter is true if the element is found; otherwise, it is false.
func Find[T comparable](l *Linked[T], v T) (*Element[T], bool) {
	for e, i := l.root.next, 0; i < l.len; e, i = e.next, i+1 {
		if e.Value == v {
			return e, true
		}
	}
	return nil, false
}

// Index returns the index of the first element with value v, or -1 if not present.
func Index[T comparable](l *Linked[T], v T) int {
	for e, i := l.root.next, 0; i < l.len; e, i = e.next, i+1 {
		if e.Value == v {
			return i
		}
	}
	return -1
}
