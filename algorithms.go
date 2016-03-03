// algorithms implements a number of common algorithm implementations
// such as quicksort, mergesort, lower bound, upper bound, binary search, etc.
package algorithms

//package main

import (
	"fmt"
	"math/rand"
)

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

func (s SortableSequence) SubRange(i, j int) *Sortable {
	return s[i:j]
}

func (s SortableSequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	seq1 := SortableSequence{3, 7, 6, 5, 4, 2, 1}
	seq2 := SortableSequence{3, 7, 6, 5, 4, 2, 1}
	fmt.Printf("starting: %v\n", seq1)
	Quicksort(seq1)
	fmt.Printf("result: %v\n", seq1)

	fmt.Println("\nTrying regular partition\n")
	fmt.Printf("starting: %v\n", seq2)
	StaticQuicksort(seq2)
	fmt.Printf("result: %v\n", seq2)
}

// Sortable is an interface for data that is sortable. It works by defining
// functions for getting the data length, determining the order of two data
// items given their indices, and a swap functiont that swaps the items
// at the given indices.
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Greater(i, j int) bool
	Equal(i, j int) bool
	SubRange(i, j int) *Sortable
	Swap(i, j int)
}

// isPartitioned checks that all elements less than the index specified by p
// are less than or equal to v[p] and all elements greater than the index p
// are greater than or equal to v[p]
func isPartitioned(v []int, p int) bool {
	for i := 0; i < len(v); i++ {
		if v[i] > p && i < p {
			return false
		}
		if v[i] < p && i > p {
			return false
		}
	}
	return true
}

// partition returns the index of the partition value such that all elements
// to the left of the index are less than the partition value and all
// elements to the right or greater than the partitioning value
func partition(v []int) int {
	p := v[0]
	i := 0
	j := len(v) - 1

	for {
		for v[i] < p {
			i++
		}
		for v[j] > p {
			j--
		}
		if i >= j {
			return j
		}
		if v[i] == v[j] {
			j--
		}
		v[i], v[j] = v[j], v[i]
	}
	return j
}

// partition partitions a given data set that implements the sortable
// interface. The initial element is chosen as the partition value
func sortablePartition(s Sortable) int {
	p := 0
	i := 0
	j := s.Len() - 1
	for {
		for s.Less(i, p) {
			i++
		}
		for s.Greater(j, p) {
			j--
		}
		if i >= j {
			return j
		}
		if s.Equal(i, j) {
			j--
		}
		if i == p {
			p = j
		} else if j == p {
			p = i
		}
		s.Swap(i, j)
	}
	return j
}

// similar to partition but instead of choosing the right most element as
// the partition value, this function choose a value at random to be the
// partitioning value
func randPartition(v []int) int {
	// select a random value as theh pivot
	p := rand.Intn(len(v))
	i := 0
	j := len(v) - 1

	for {
		for v[i] < p {
			i++
		}
		for v[j] > p {
			j--
		}
		if i >= j {
			return j
		}
		if v[i] == v[j] {
			j--
		}
		v[i], v[j] = v[j], v[i]
	}
}

func sortableRandPartition(s Sortable) int {
	p := rand.Intn(s.Len())
	i := 0
	j := s.Len() - 1

	for {
		for s.Less(i, p) {
			i++
		}
		for s.Less(j, p) {
			j--
		}
		if i >= j {
			return j
		}
		if !s.Equal(i, j) {
			j--
		}
		s.Swap(i, j)
	}
	return j
}

// Quicksort implementation using parition
func Quicksort(s Sortable) {
	p := sortablePartition(s)
	if s.Len() <= 2 {
		return
	}
	fmt.Printf("p: %d\nv: %v\n", p, s)
	fmt.Println("Starting left recursive call...")
	if p > 0 {
		Quicksort(s)
	}
	fmt.Println("Starting right recursive call...")
	if p < s.Len()-1 {
		Quicksort(s.SubRange(p, s.Len()-1))
	}
}

// StaticQuicksort is quicksort using the first element as the partition
func StaticQuicksort(v []int) {
	p := partition(v)
	if len(v) == 2 {
		return
	}
	if p > 0 {
		StaticQuicksort(v[:p])
	}
	if p < len(v)-1 {
		StaticQuicksort(v[p+1:])
	}
}
