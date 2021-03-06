package main

import "fmt"

func makeMinHeap(slice []int, i int, size int) {
	// return back if leaf nodes
	if i > size/2-1 {
		return
	}

	lchild := i*2 + 1
	rchild := i*2 + 2

	// return if lchild or rchild index is out of array range
	if lchild >= size || rchild >= size {
		return
	}

	if slice[lchild] < slice[rchild] && slice[lchild] < slice[i] {
		slice[lchild], slice[i] = slice[i], slice[lchild]
		// maintain order of sub tree as well
		makeMinHeap(slice, lchild, size)
	} else if slice[rchild] < slice[lchild] && slice[rchild] < slice[i] {
		slice[rchild], slice[i] = slice[i], slice[rchild]
		// maintain order of sub tree as well
		makeMinHeap(slice, rchild, size)
	}
}

func minHeap(slice []int, size int) {
	for i := size / 2; i >= 0; i-- {
		makeMinHeap(slice, i, size)
	}
}

func deleteElement(slice *[]int, index int) {
	size := len(*slice)
	// Add last element at the index
	(*slice)[index] = (*slice)[size-1]
	// Remove the last element
	*slice = (*slice)[:size-1]

	makeMinHeap((*slice), index, size-1)
}

func main() {
	heap := []int{5, 3, 17, 10, 84, 19, 6, 22, 9, 18, 20, 1, 24}
	size := len(heap)

	fmt.Println(heap)
	minHeap(heap, size)
	fmt.Println(heap)

	fmt.Println()

	fmt.Println(heap)
	deleteElement(&heap, 0)
	fmt.Println(heap)
}
