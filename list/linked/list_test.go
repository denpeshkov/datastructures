// Based on https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/container/list/list_test.go
package linked

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New[any]()
	checkListPointers(t, l, []*Element[any]{})

	// Single element list
	e := l.InsertFront("a")
	checkListPointers(t, l, []*Element[any]{e})
	l.Remove(e)
	checkListPointers(t, l, []*Element[any]{})

	// Bigger list
	e2 := l.InsertFront(2)
	e1 := l.InsertFront(1)
	e3 := l.InsertBack(3)
	e4 := l.InsertBack("banana")
	checkListPointers(t, l, []*Element[any]{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*Element[any]{e1, e3, e4})
	l.Remove(e4)
	l.Remove(e3)
	e4 = l.InsertBack(e4.Value)
	e3 = l.InsertBack(e3.Value)
	checkListPointers(t, l, []*Element[any]{e1, e4, e3})

	e2 = l.InsertBefore(2, e1) // insert before front
	checkListPointers(t, l, []*Element[any]{e2, e1, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e4) // insert before middle
	checkListPointers(t, l, []*Element[any]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e3) // insert before back
	checkListPointers(t, l, []*Element[any]{e1, e4, e2, e3})
	l.Remove(e2)

	e2 = l.InsertAfter(2, e1) // insert after front
	checkListPointers(t, l, []*Element[any]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e4) // insert after middle
	checkListPointers(t, l, []*Element[any]{e1, e4, e2, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e3) // insert after back
	checkListPointers(t, l, []*Element[any]{e1, e4, e3, e2})
	l.Remove(e2)

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			sum += i
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *Element[any]
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkListPointers(t, l, []*Element[any]{})
}

func TestRemove(t *testing.T) {
	l := New[any]()
	e1 := l.InsertBack(1)
	e2 := l.InsertBack(2)
	checkListPointers(t, l, []*Element[any]{e1, e2})
	e := l.Front()
	l.Remove(e)
	checkListPointers(t, l, []*Element[any]{e2})
	l.Remove(e)
	checkListPointers(t, l, []*Element[any]{e2})
}

func TestInsertElementFromDifferentList(t *testing.T) {
	l1 := New[any]()
	l1.InsertBack(1)
	l1.InsertBack(2)

	l2 := New[any]()
	l2.InsertBack(3)
	l2.InsertBack(4)

	e := l1.Front()
	l2.Remove(e) // l2 should not change because e is not an element of l2
	if n := l2.Len(); n != 2 {
		t.Errorf("l2.Len() = %d, want 2", n)
	}

	l1.InsertBefore(8, e)
	if n := l1.Len(); n != 3 {
		t.Errorf("l1.Len() = %d, want 3", n)
	}
}

func TestRemovedElementIsUnlinked(t *testing.T) {
	l := New[any]()
	l.InsertBack(1)
	l.InsertBack(2)

	e := l.Front()
	l.Remove(e)
	if e.Value != 1 {
		t.Errorf("e.value = %d, want 1", e.Value)
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.Prev() != nil {
		t.Errorf("e.Prev() != nil")
	}
}

// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
func TestInsertBeforeUnknownMark(t *testing.T) {
	l := New[any]()
	l.InsertBack(1)
	l.InsertBack(2)
	l.InsertBack(3)
	l.InsertBefore(1, new(Element[any]))
	checkList(t, l, []any{1, 2, 3})
}

// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
func TestInsertAfterUnknownMark(t *testing.T) {
	l := New[any]()
	l.InsertBack(1)
	l.InsertBack(2)
	l.InsertBack(3)
	l.InsertAfter(1, new(Element[any]))
	checkList(t, l, []any{1, 2, 3})
}

/* func TestIterator(t *testing.T) {
	l := NewLinked[any]()

	it := l.Iter()

} */

func checkListLen[T any](t *testing.T, l *Linked[T], len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers[T any](t *testing.T, l *Linked[T], es []*Element[T]) {
	root := &l.root

	if !checkListLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 {
		if l.root.next != nil && l.root.next != root || l.root.prev != nil && l.root.prev != root {
			t.Errorf("l.root.next = %p, l.root.prev = %p; both should be nil or %p", l.root.next, l.root.prev, root)
		}
		return
	}

	// check internal and external prev/next connections
	for i, e := range es {
		prev := root
		Prev := (*Element[T])(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := e.prev; p != prev {
			t.Errorf("elt[%d](%p).prev = %p, want %p", i, e, p, prev)
		}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := root
		Next := (*Element[T])(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := e.next; n != next {
			t.Errorf("elt[%d](%p).next = %p, want %p", i, e, n, next)
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func checkList(t *testing.T, l *Linked[any], es []any) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}
