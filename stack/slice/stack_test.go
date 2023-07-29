package slice_test

import (
	"testing"

	"github.com/denpeshkov/datastructures/stack/internal"
	. "github.com/denpeshkov/datastructures/stack/slice"
)

var gen = func() internal.Stack[int] { return new(Stack[int]) }

func TestLen(t *testing.T) {
	internal.TestLen(t, gen)
}

func TestEmpty(t *testing.T) {
	internal.TestEmpty(t, gen)
}

func TestPush(t *testing.T) {
	internal.TestPush(t, gen)
}

func TestPop(t *testing.T) {
	internal.TestPop(t, gen)
}

func TestPeek(t *testing.T) {
	internal.TestPeek(t, gen)
}
