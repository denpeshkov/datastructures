package internal

import (
	"testing"
)

type Gen[T any] func() Stack[T]

type Stack[T any] interface {
	Push(x T)
	Pop() T
	Peek() T
	Len() int
	Empty() bool
}

func generate[T any](g func() Stack[T], e ...T) Stack[T] {
	s := g()
	for _, v := range e {
		s.Push(v)
	}

	return s
}

func testLen[T any](t *testing.T, s Stack[T], l int) {
	if s.Len() != l {
		t.Errorf("expected len to be: %v; got: %v", l, s.Len())
	}
}

func TestEmpty(t *testing.T, g func() Stack[int]) {
	s := g()

	if !s.Empty() {
		t.Errorf("expected empty stack; got %v", s)
	}
	testLen(t, s, 0)

	for i := 0; i < 5; i++ {
		s.Push(i)

		if s.Empty() {
			t.Errorf("expected not empty stack; got %v", s)
			testLen(t, s, i+1)
		}
	}
}

func TestLen(t *testing.T, g func() Stack[int]) {
	tests := []struct {
		s   Stack[int]
		len int
	}{
		{generate(g), 0},
		{generate(g, 1), 1},
		{generate(g, 1, 2, 3), 3},
		{generate(g, 1, 2, 3, 4), 4},
		{generate(g, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 10},
	}

	for _, test := range tests {
		testLen(t, test.s, test.len)
	}
}

func TestPeek(t *testing.T, g func() Stack[int]) {
	tests := []struct {
		s         Stack[int]
		expectedV int
	}{
		{generate(g), 0},
		{generate(g, 1), 1},
		{generate(g, -1), -1},
		{generate(g, 1, 2), 2},
		{generate(g, 2, 4, 6, 0, 0, -1), -1},
		{generate(g, 1, 1, 1, 1, 1, 2), 2},
	}

	for _, test := range tests {
		v := test.s.Peek()

		if v != test.expectedV {
			t.Errorf("expected v to be: %v; got: %v", test.expectedV, v)
		}
	}
}

func TestPop(t *testing.T, g func() Stack[int]) {
	tests := []struct {
		s             Stack[int]
		expectedV     int
		expectedPeekV int
	}{
		{generate(g), 0, 0},
		{generate(g, 1), 1, 0},
		{generate(g, -1), -1, 0},
		{generate(g, 1, 2), 2, 1},
		{generate(g, 2, 4, 6, 0, 0, -1), -1, 0},
		{generate(g, 1, 1, 1, 1, 1, 2), 2, 1},
	}

	for _, test := range tests {
		v := test.s.Pop()
		peekV := test.s.Peek()

		if v != test.expectedV {
			t.Errorf("expected v to be: %v; got: %v", test.expectedV, v)
		}
		if peekV != test.expectedPeekV {
			t.Errorf("expected top value after Pop() to be: %v; got: %v", test.expectedPeekV, peekV)
		}
	}
}

func TestPush(t *testing.T, g func() Stack[int]) {
	tests := []struct {
		s         Stack[int]
		expectedV int
	}{
		{generate(g), 0},
		{generate(g), 2},
		{generate(g, 1), 1},
		{generate(g, 1), 4},
		{generate(g, -1), 10},
		{generate(g, 1, 2), 2},
		{generate(g, 1, 2), -100},
		{generate(g, 2, 4, 6, 0, 0, -1), -1},
		{generate(g, 2, 4, 6, 0, 0, -1), 3},
		{generate(g, 1, 1, 1, 1, 1, 2), 11},
	}

	for _, test := range tests {
		l := test.s.Len()
		test.s.Push(test.expectedV)

		v := test.s.Peek()

		if v != test.expectedV {
			t.Errorf("expected v to be: %v; got: %v", test.expectedV, v)
		}
		testLen(t, test.s, l+1)
	}
}
