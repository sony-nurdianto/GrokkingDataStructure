package main

import (
	"fmt"
)

type Heap[T any] struct {
	data     []T
	lessFunc func(a, b T) bool // Fungsi pembanding untuk mendukung min-heap dan max-heap
}

// NewHeap membuat heap dengan fungsi pembanding
func NewHeap[T any](lessFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{data: []T{}, lessFunc: lessFunc}
}

func (h *Heap[T]) heapifyUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.lessFunc(h.data[parent], h.data[index]) {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap[T]) heapifyDown(index int) {
	n := len(h.data)
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < n && h.lessFunc(h.data[largest], h.data[left]) {
		largest = left
	}
	if right < n && h.lessFunc(h.data[largest], h.data[right]) {
		largest = right
	}

	if largest != index {
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		h.heapifyDown(largest)
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return top, true
}

func (h *Heap[T]) Peek() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap[T]) Display() {
	fmt.Println(h.data)
}

func main() {
	// Max-Heap (default, seperti kode Anda sebelumnya)
	maxHeap := NewHeap(func(a, b int) bool { return a < b })
	maxHeap.Push(10)
	maxHeap.Push(3)
	maxHeap.Push(20)

	fmt.Println("Max Heap:")
	maxHeap.Display()

	// Min-Heap
	minHeap := NewHeap(func(a, b int) bool { return a > b })
	minHeap.Push(10)
	minHeap.Push(3)
	minHeap.Push(20)

	fmt.Println("Min Heap:")
	minHeap.Display()
}
