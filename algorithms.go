// algorithms implements a number of common algorithm implementations
// such as quicksort, mergesort, lower bound, upper bound, binary search, etc.
package algorithms

import (
	"fmt"
	"math/rand"
)

func main() {
	test := []int{3, 7, 6, 5, 4, 2, 1}
	fmt.Printf("test1: %v\n", test)
	p := randPartition(test)
	fmt.Printf("result1: %v\n", test)
	if isPartitioned(test, p) {
		fmt.Printf("test1 is partitioned at %d\n", p)
	} else {
		fmt.Println("test1 is not partitioned")
	}

	test2 := []int{4, 8, 7, 6, 5, 3, 2, 1}
	fmt.Printf("test2: %v\n", test2)
	p = randPartition(test2)
	fmt.Printf("result2: %v\n", test2)
	if isPartitioned(test2, p) {
		fmt.Printf("test2 is partitioned at %d\n", p)
	} else {
		fmt.Println("test2 is not partitioned")
	}
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
		v[i], v[j] = v[j], v[i]
	}
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
		v[i], v[j] = v[j], v[i]
	}
}

// Quicksort implementation using parition
func Quicksort(v []int) {

}
