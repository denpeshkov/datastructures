// Package unionfind contains an implementation of a a disjoint-set (union-find) data structure.
package unionfind

type item struct {
	p    int
	size int
}

// TODO
// UnionFind represent an union-find data structure.
type UnionFind struct {
	s []item
}

func New(size int) *UnionFind {
	uf := &UnionFind{}

	uf.s = make([]item, size)
	for i := 0; i < size; i++ {
		uf.s[i] = item{i, 1}
	}

	return uf
}

func (uf *UnionFind) Union(p, q int) {
	s := uf.s
	i, j := uf.Find(p), uf.Find(q)

	if i == j {
		return
	}

	if s[i].size <= s[j].size {
		s[i].p = j
		s[j].size += s[i].size
	} else {
		s[j].p = i
		s[i].size += s[j].size
	}
}

func (uf *UnionFind) Find(i int) int {
	s := uf.s
	for s[i].p != i {
		i = s[i].p
	}
	return i
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}
