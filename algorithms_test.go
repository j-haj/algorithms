package algorithms

import (
	"testing"
)

// ----------------------------------------------------------------------------
// Sortable Interface
// ----------------------------------------------------------------------------
type SortableSequence []int

func (s SortableSequence) Len() int {
	return len(s)
}

func (s SortableSequence) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s SortableSequence) Greater(i, j int) bool {
	if s[i] > s[j] {
		return true
	}
	return false
}

func (s SortableSequence) Equal(i, j int) bool {
	if s[i] == s[j] {
		return true
	}
	return false
}

func (s SortableSequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isSorted(s Sortable) bool {
	for i := 0; i < s.Len()-2; i++ {
		if !s.Less(i, i+1) && !s.Equal(i, i+1) {
			return false
		}
	}
	return true
}

var seq1 = SortableSequence{3, 7, 6, 5, 4, 2, 1}
var srt1 = SortableSequence{1, 2, 3, 4, 5, 6, 7}

var seq2 = SortableSequence{9, 8, 7, 6, 5, 4, 3, 2, 1}
var srt2 = SortableSequence{1, 2, 3, 4, 5, 6, 7, 8, 9}

var seq3 = SortableSequence{8, 8, 7, 7, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1}
var srt3 = SortableSequence{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8}

// ----------------------------------------------------------------------------
// Partition Tests
// ----------------------------------------------------------------------------
func TestPartition1(t *testing.T) {
	test := []int{3, 7, 6, 5, 4, 2, 1}
	expected := []int{1, 2, 3, 5, 4, 6, 7}
	partition(test)
	for i := 0; i < len(test); i++ {
		if test[i] != expected[i] {
			t.Fatalf("expected %d == %d\n\tactual %d != %d, at index %d",
				expected[i], expected[i], test[i], expected[i], i)
		}
	}
}

func TestPartition2(t *testing.T) {
	test := []int{4, 8, 7, 6, 5, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	partition(test)
	for i := 0; i < len(test); i++ {
		if test[i] != expected[i] {
			t.Fatalf("expected %d == %d\nactual %d != %d, at index %d",
				expected[i], expected[i], test[i], expected[i], i)
		}
	}
}

func TestRandPartition1(t *testing.T) {
	test := []int{3, 7, 6, 5, 4, 2, 1}
	expected := []int{3, 1, 2, 5, 4, 6, 7}
	p := randPartition(test)
	if p != 5 {
		t.Fatalf("expected p = 5, got p  = %d\n", p)
	}
	for i := 0; i < len(test); i++ {
		if test[i] != expected[i] {
			t.Fatalf("expected test[%d] = %d, got test[%d] = %d\n",
				i, expected[i], i, test[i])
		}
	}
}

// ----------------------------------------------------------------------------
// Sort Tests
// ----------------------------------------------------------------------------
func TestStaticQuicksort1(t *testing.T) {
	test := []int{6, 5, 1, 3, 9, 2, 8, 6, 6, 5}
	expected := []int{1, 2, 3, 5, 5, 6, 6, 6, 8, 9}
	StaticQuicksort(test)
	for i := 0; i < len(test); i++ {
		if test[i] != expected[i] {
			t.Fatalf("expected test[%d] = %d; got test[%d] = %d\n", i, test[i], i, expected[i])
		}
	}
}

func TestStaticQuicksort2(t *testing.T) {
	test := []int{6, 6, 5, 5, 1, 1, 3, 3, 9, 9, 2, 2, 8, 8, 6, 6, 6, 5, 5}
	expected := []int{1, 1, 2, 2, 3, 3, 5, 5, 5, 5, 6, 6, 6, 6, 6, 8, 8, 9, 9}
	StaticQuicksort(test)
	for i := 0; i < len(test); i++ {
		if test[i] != expected[i] {
			t.Fatalf("expected test[%d] = %d; got test[%d] = %d\n", i, test[i], i, expected[i])
		}
	}
}

func TestQuicksort1(t *testing.T) {
	Quicksort(seq1, 0, seq1.Len()-1, 0)
	if !isSorted(seq1) {
		t.Fatalf("expected seq1 = %v; got seq1 = %v.", seq1, srt1)
	}
}

func TestQuicksort2(t *testing.T) {
	Quicksort(seq2, 0, seq2.Len()-1, 0)
	if !isSorted(seq2) {
		t.Fatalf("expected seq2 = %v; got seq2 = %v", seq2, srt2)
	}
}

func TestQuicksort3(t *testing.T) {
	Quicksort(seq3, 0, seq3.Len()-1, 0)
	if !isSorted(seq3) {
		t.Fatalf("expected seq3 = %v; got seq3 = %v.", seq3, srt3)
	}
}
