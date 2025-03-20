package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	data []T
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{data: []T{}}
}

func (h *Heap[T]) heapifyUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.data[index] > h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *Heap[T]) Enqueu(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap[T]) heapifyDown(index int) {
	n := len(h.data)
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < n && h.data[left] > h.data[largest] {
		largest = left
	}

	if right < n && h.data[right] > h.data[largest] {
		largest = right
	}

	if largest != index {
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		h.heapifyDown(largest)
	}
}

func (h *Heap[T]) Dequeu() (T, bool) {
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

func (h *Heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap[T]) Display() {
	fmt.Println(h.data)
}

func main() {
	heap := NewHeap[int]()
	heap.Enqueu(10)
	heap.Enqueu(3)
	heap.Enqueu(20)
	heap.Enqueu(50)

	heap.Display()
}
