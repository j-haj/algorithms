package heap_test

import (
	"github.com/j-haj/algorithms-data-structures/heap"
	"testing"
)

func TestHeapifyMax(t *testing.T) {
	expectedHeap := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	testHeap := []int{14, 16, 10, 8, 7, 9, 3, 2, 4, 1}
	resultHeap := heap.HeapifyMax(testHeap, 0)
	for i := 0; i < len(resultHeap); i++ {
		if resultHeap[i] != expectedHeap[i] {
			t.Error("HeapifyMax failed.")
		}
	}

}

func TestHeapifyMin(t *testing.T) {
	expectedHeap := []int{1, 2, 3, 4, 5, 6, 7}
	testHeap := []int{3, 2, 1, 4, 5, 6, 7}
	resultHeap := heap.HeapifyMin(testHeap, 0)
	for i := 0; i < len(resultHeap); i++ {
		if expectedHeap[i] != resultHeap[i] {
			t.Error("HeapifyMin failed.")
		}
	}
}

func TestCreateMaxHeap(t *testing.T) {
	initialElements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	resultHeap := heap.CreateMaxHeap(initialElements)
	if !heap.IsMaxHeap(resultHeap) {
		t.Error("CreateMaxHeap failed.")
	}
}

func TestCreateMinHeap(t *testing.T) {
	initialElements := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	resultHeap := heap.CreateMinHeap(initialElements)
	if !heap.IsMinHeap(resultHeap) {
		t.Error("CreateMinHeap failed.")
	}
}

func TestHeapsortMin(t *testing.T) {
	test1 := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	test2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	res1 := heap.Heapsort(test1, true)
	res2 := heap.Heapsort(test2, false)

	for i := 0; i < len(test1); i++ {
		if res1[i] != test2[i] {
			t.Error("Heapsort min to max failed.")
		}
		if res2[i] != test1[i] {
			t.Error("Heapsort max to min failed")
		}
	}
}
