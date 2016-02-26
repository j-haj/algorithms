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
	partition(test)
	fmt.Printf("result1: %v\n", test)

	test2 := []int{4, 8, 7, 6, 5, 3, 2, 1}
	fmt.Printf("test2: %v\n", test2)
	partition(test2)
	fmt.Printf("result2: %v\n", test2)
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
		fmt.Printf("(i, j) --> (%d, %d)\n", i, j)
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
