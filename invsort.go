// Package perm handles manipulation of permutations.
package perm

import "sort"

type invSorter struct {
	coll sort.Interface // the collection to be sorted

	// perm is the forward permutation if coll; perm[i] is the current offset of
	// the ith element of the original collection.
	perm []int

	// inv is the inverse permutation of coll; inv[i] is the original offset of
	// the element now at i.
	inv []int
}

func (r *invSorter) Len() int           { return len(r.perm) }
func (r *invSorter) Less(i, j int) bool { return r.coll.Less(i, j) }
func (r *invSorter) Swap(i, j int) {
	r.coll.Swap(i, j)
	r.perm[i], r.perm[j] = r.perm[j], r.perm[i]
	r.inv[r.perm[i]] = i
	r.inv[r.perm[j]] = j
}

// SortInverse sorts the specified collection and returns a slice of integer
// offsets of length coll.Len() giving the reverse permutation.
func SortInverse(coll sort.Interface) []int {
	n := coll.Len()
	r := &invSorter{
		coll: coll,
		perm: make([]int, n),
		inv:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		r.perm[i] = i
		r.inv[i] = i
	}
	sort.Sort(r)
	return r.inv
}
