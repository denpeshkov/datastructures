// Package bitset contains the implementation of a set backed by a bit array.
package bitset

import (
	"fmt"
	"math/bits"
	"slices"
	"strings"
)

const (
	wordSize   = 64
	lgWordSize = 6
)

// wordInd returns the index of the word that should contain bit e.
func wordInd(e uint) int {
	return int(e >> lgWordSize) // e/64
}

// bitInd returns the index of the bit e inside the word.
func bitInd(e uint) int {
	return int(e & (wordSize - 1)) // e%64
}

/*
Set is an implementation of a set ADT backed by a bit array.
It can store elements in the range [0..[math.MaxUint64]].
Default value represents an empty set and is ready to use.
*/
type Set struct {
	set []uint64
	// number of bits
	len uint // TODO maybe delete
}

// New creates a new set with a hint that at least length bits are required.
// If length is 0, returns an empty set.
func New(length uint) *Set {
	if length == 0 {
		return &Set{}
	}
	return &Set{set: make([]uint64, wordInd(length-1)+1), len: length}
}

// From creates a new set with elements from s.
func From(s []uint) *Set {
	rs := &Set{}

	for _, e := range s {
		rs.Add(e)
	}
	return rs
}

// Values returns all the elements in the set.
func (b *Set) Values() []uint {
	res := make([]uint, 0, b.Size())

	for wI, w := range b.set {
		for i := uint(0); i < wordSize; i++ {
			if w&(1<<i) != 0 {
				res = append(res, wordSize*uint(wI)+i)
			}
		}
	}
	return res
}

// AsBits returns the contents of a set as a string of bits.
func (b *Set) AsBits() string {
	var sb strings.Builder
	for _, w := range b.set {
		fmt.Fprintf(&sb, "%064b.", w)
	}
	return sb.String()
}

// Union returns the union of two sets. Specifically, resulting in a set where any bit set to 1 in either of the original sets is set to 1 in the returned set.
func (s *Set) Union(other *Set) *Set {
	smS, lgS := s, other
	if len(smS.set) > len(lgS.set) {
		smS, lgS = lgS, smS
	}

	res := lgS.Clone()
	for i, w := range smS.set {
		res.set[i] |= w
	}
	return res
}

// Difference returns the difference of two sets.
// Specifically, the bit in the resulting set is set to 1 if it is set to 1 in this set and set to 0 in other set.
func (s *Set) Difference(other *Set) *Set {
	res := s.Clone()

	for i := uint(0); i < min(s.len, other.len); i++ {
		res.set[i] &^= other.set[i]
	}
	return res
}

// Add adds element e to the set. Specifically, sets bit at index e to 1.
func (s *Set) Add(e uint) {
	if e >= s.len {
		s.grow(e)
	}

	wI, bI := wordInd(e), bitInd(e)
	s.set[wI] |= 1 << bI
}

// Contains returns true if this set contains the specified element.
// Specifically, it returns true if the bit at index e is set to 1. If e is larger than [Set.Len] returns false.
func (s *Set) Contains(e uint) bool {
	wI, bI := wordInd(e), bitInd(e)
	if e > s.len {
		return false
	}
	return s.set[wI]&(1<<bI) != 0
}

// Remove removes the element e from the set. Specifically, it sets the bit at index e to 0.
// If e is larger than [Set.Len] set is left unchanged.
func (s *Set) Remove(e uint) {
	if e >= s.len {
		return
	}
	wI, bI := wordInd(e), bitInd(e)
	s.set[wI] &^= 1 << bI
}

// RemoveAll removes all the elements from the set. Specifically, it sets all the bits to 0.
func (s *Set) RemoveAll() {
	for i := range s.set {
		s.set[i] = 0
	}
}

// Clone clones this set.
func (s *Set) Clone() *Set {
	return &Set{set: slices.Clone(s.set), len: s.len}
}

// Equal returns true if the provided set and this set contain the same elements.
// Namely, the lengths of the sets can be different.
func (s *Set) Equal(other *Set) bool {
	if s == other {
		return true
	}

	for i, s1, s2 := 0, s.set, other.set; i < max(len(s1), len(s2)); i++ {
		// when sets have different sizes
		if (i >= len(s1) && s2[i] != 0) || (i >= len(s2) && s1[i] != 0) {
			return false
		}
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// Size returns the number of the elements in the set.
// Specifically, it returns the number of bits set to 1 in a bitset (population count).
func (s *Set) Size() int {
	cnt := 0
	for _, w := range s.set {
		cnt += bits.OnesCount64(w)
	}
	return cnt
}

// grow extends the bitset if necessary to ensure sufficient space for storing an element with index i.
func (s *Set) grow(i uint) {
	wrdCnt := wordInd(i) + 1
	s.set = slices.Grow(s.set, wrdCnt-len(s.set))[:wrdCnt]
	s.len = i + 1
}
