// algorithms implements a number of common algorithm implementations
// such as quicksort, mergesort, lower bound, upper bound, binary search, etc.
package algorithms

import (
	"fmt"
	"math/rand"
)

// Sortable is an interface for data that is sortable. It works by defining
// functions for getting the data length, determining the order of two data
// items given their indices, and a swap functiont that swaps the items
// at the given indices.
type Sortable interface {
	Len() int
	Less(i, j int) bool
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

// Quicksort implementation using parition
func Quicksort(v []int) {
	p := partition(v)
	if len(v) == 2 {
		return
	}
	fmt.Printf("p: %d\nv: %v\n", p, v)
	fmt.Println("Starting left recursive call...")
	if p > 0 {
		Quicksort(v[:p])
	}
	fmt.Println("Starting right recursive call...")
	if p < len(v)-1 {
		Quicksort(v[p+1:])
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
