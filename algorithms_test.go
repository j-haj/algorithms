package algorithms

import (
	"testing"
)

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
