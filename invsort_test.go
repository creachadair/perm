package perm

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortInverse(t *testing.T) {
	input := []string{
		"oscar", "romeo", "delta", "echo", "romeo", "india", "november", "golf",
	}
	t.Logf("Input:   %+q", input)

	// Sort the collection and make sure it is correctly sorted.
	work := make([]string, len(input))
	copy(work, input)
	inv := SortInverse(sort.StringSlice(work))
	if !sort.IsSorted(sort.StringSlice(work)) {
		t.Errorf("SortInverse did not work, result is out of order:\n%+q", work)
	}
	t.Logf("Sorted:  %+q", work)
	t.Logf("Inverse: %v", inv)

	// Apply the reverse permutation and verify it restores the input order.
	got := make([]string, len(inv))
	for i, v := range inv {
		got[i] = work[v]
	}
	if !reflect.DeepEqual(got, input) {
		t.Errorf("SortInverse: inverse is wrong:\ngot  %+q\nwant %+q", got, input)
	}
}
