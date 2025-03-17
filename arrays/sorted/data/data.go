package data

import (
	"errors"
	"fmt"
)

type SortedArray struct {
	array   []int
	maxSize int
	size    int
}

func NewSortedArray(maxSize int) *SortedArray {
	return &SortedArray{
		array:   make([]int, maxSize),
		maxSize: maxSize,
		size:    0,
	}
}

func (sa *SortedArray) Insert(value int) error {
	if sa.size >= sa.maxSize {
		return errors.New("insert: array is full")
	}

	i := sa.size

	for i > 0 && sa.array[i-1] > value {
		sa.array[i] = sa.array[i-1]
		i--
	}

	sa.array[i] = value
	sa.size++

	return nil
}

func (sa *SortedArray) Search(target int) int {
	low, high := 0, sa.size-1

	for low <= high {
		mid := low + (high-low)/2

		if sa.array[mid] == target {
			return mid // Ketemu
		} else if sa.array[mid] < target {
			low = mid + 1 // Cari di kanan
		} else {
			high = mid - 1 // Cari di kiri
		}
	}

	return -1 // Tidak ketemu
}

func (sa *SortedArray) Delete(target int) error {
	index := sa.Search(target)
	if index == -1 {
		return fmt.Errorf("delete: target %d not found", target)
	}

	fmt.Println(index)

	for i := index; i < sa.size-1; i++ {
		sa.array[i] = sa.array[i+1]
	}

	sa.array[sa.size-1] = 0

	sa.size--
	return nil
}

func (sa *SortedArray) DeleteByIndex(index int) error {
	if sa.size < 0 {
		return errors.New("delete index: array is empty")
	}

	if index > sa.maxSize {
		return errors.New("delete index: index offset")
	}

	for i := index; i < sa.size-1; i++ {
		sa.array[i] = sa.array[i+1]
	}

	sa.array[sa.size-1] = 0

	sa.size--
	return nil
}
