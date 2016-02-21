// Heap contains functions for creating min and max heaps as well as an implementation of heapsort
package heap

// Heapsort is an implementation of the heapsort algorithm. Set ascending to true
// if the passed elements should be sorted smallest to largest and set ascending
// to false if the elements should be sorted largest to smallest
func Heapsort(elements []int, ascending bool) []int {
	N := len(elements)
	sortedElements := make([]int, N)
	heap := make([]int, N)

	if ascending {
		heap = CreateMinHeap(elements)

	} else {
		heap = CreateMaxHeap(elements)
	}

	index := 0
	for len(heap) > 0 {
		sortedElements[index] = heap[0]
		if ascending {
			heap = CreateMinHeap(heap[1:])
		} else {
			heap = CreateMaxHeap(heap[1:])
		}
		index++
	}
	return sortedElements
}

func leftIndex(index int) int {
	return index<<1 + 1
}

func rightIndex(index int) int {
	return (index + 1) << 1
}

// Heapify the given max-heap at the specified index
func HeapifyMax(heap []int, index int) []int {
	leftIndex := leftIndex(index)
	rightIndex := rightIndex(index)
	largestIndex := index

	// Check left child
	if (leftIndex < len(heap)) && (heap[leftIndex] > heap[index]) {
		largestIndex = leftIndex
	}

	// Check right child
	if (rightIndex < len(heap)) && (heap[rightIndex] > heap[largestIndex]) {
		largestIndex = rightIndex
	}

	if largestIndex != index {
		// Swap elements and make recursive call
		heap[largestIndex], heap[index] = heap[index], heap[largestIndex]
		return HeapifyMax(heap, largestIndex)
	} else {
		return heap
	}
}

// Heapify the given min-heap at the specified index
func HeapifyMin(heap []int, index int) []int {
	leftIndex := leftIndex(index)
	rightIndex := rightIndex(index)
	minIndex := index

	// Check left child
	if leftIndex < len(heap) && heap[leftIndex] < heap[index] {
		minIndex = leftIndex
	}
	// Check right child
	if rightIndex < len(heap) && heap[rightIndex] < heap[minIndex] {
		minIndex = rightIndex
	}

	if minIndex != index {
		heap[minIndex], heap[index] = heap[index], heap[minIndex]
		return HeapifyMin(heap, minIndex)
	} else {
		return heap
	}
}

// Create a max-heap (O(lg(n)))
func CreateMaxHeap(elements []int) []int {
	heap := make([]int, len(elements))
	copy(heap, elements)
	for i := int(len(elements)/2) + 1; i >= 0; i-- {
		heap = HeapifyMax(heap, i)
	}
	return heap
}

// Create a min-heap (O(lg(n)))
func CreateMinHeap(elements []int) []int {
	heap := make([]int, len(elements))
	copy(heap, elements)
	for i := int(len(elements)/2) + 1; i >= 0; i-- {
		heap = HeapifyMin(heap, i)
	}
	return heap
}

func IsMaxHeap(heap []int) bool {
	for i := 0; i < int(len(heap)/2)+1; i++ {
		if leftIndex(i) > len(heap)-1 {
			return true
		}
		if heap[i] < heap[leftIndex(i)] {
			return false
		}
		if heap[i] < heap[rightIndex(i)] {
			return false
		}
	}
	return true
}

func IsMinHeap(heap []int) bool {
	for i := 0; i < int(len(heap)/2)+1; i++ {
		if leftIndex(i) > len(heap)-1 {
			return true
		}
		if heap[i] > heap[leftIndex(i)] {
			return false
		}
		if heap[i] > heap[rightIndex(i)] {
			return false
		}
	}
	return true
}
