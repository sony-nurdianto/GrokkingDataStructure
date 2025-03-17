package data

import (
	"errors"
	"fmt"
)

type UnsortedArray struct {
	Array   []int
	MaxSize int
	Size    int
}

func NewUnsortedArray(maxSize int) *UnsortedArray {
	return &UnsortedArray{
		Array:   make([]int, maxSize),
		MaxSize: maxSize,
		Size:    0,
	}
}

func (ua *UnsortedArray) Insert(value int) error {
	if ua.Size >= ua.MaxSize {
		return errors.New("insert: array is full")
	}

	ua.Array[ua.Size] = value
	ua.Size += 1

	return nil
}

func (ua *UnsortedArray) Delete(index int) error {
	if ua.Size == 0 {
		return errors.New("delete: array is empty")
	}

	if index < 0 || index >= ua.Size {
		return fmt.Errorf("index %d offset range", index)
	}

	ua.Array[index] = ua.Array[ua.Size-1]
	ua.Size -= 1

	return nil
}

func (ua *UnsortedArray) Find(target int) int {
	for _, v := range ua.Array {
		if v == target {
			return v
		}
	}

	return -1
}

func (ua *UnsortedArray) Traverse(cb func(int)) {
	for _, v := range ua.Array {
		cb(v)
	}
}
