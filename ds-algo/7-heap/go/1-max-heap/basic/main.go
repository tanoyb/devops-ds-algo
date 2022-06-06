package main

import "fmt"

//left child index= [parent index]*2 + 1
//right child index = [parent index]*2 + 2

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

//heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

//get the left child index
func left(index int) int {
	return (2 * index) + 1
}

//get the right child index
func right(index int) int {
	return (2 * index) + 2
}

//swap keys in the array
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	if len(h.array) == 0 {
		fmt.Println("empty array")
		return -1
	}
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)

	return extracted
}

//heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftChildIndex, rightChildIndex := left(index), right(index)
	childIndexToCompare := 0

	//loop while index has at least one child
	for leftChildIndex <= lastIndex {
		if leftChildIndex == lastIndex {
			//when left child is the only child
			childIndexToCompare = leftChildIndex
		} else if h.array[leftChildIndex] > h.array[rightChildIndex] {
			//when left child is larger
			childIndexToCompare = leftChildIndex
		} else {
			//when right child is larger
			childIndexToCompare = rightChildIndex
		}
		//compare array value of current index to larger child and
		//swap if smaller
		if h.array[index] < h.array[childIndexToCompare] {
			h.swap(index, childIndexToCompare)
			index = childIndexToCompare
			leftChildIndex, rightChildIndex = left(index), right(index)
		} else {
			//break
			return
		}

	}
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	keys := []int{10, 14, 18, 20, 30, 40, 45, 47, 49, 50}
	for _, v := range keys {
		m.Insert(v)
	}
	fmt.Println(m)
	m.Extract()
	fmt.Println(m)
	m.Extract()
	fmt.Println(m)
	m.Extract()
	fmt.Println(m)

}
