package main

import (
	"fmt"
)

type DynamicArray struct {
	Array    []int
	Capacity int
	Size     int
}

func NewDynamicArray(initCap int) *DynamicArray {
	initCap = max(initCap, 1) // Pastikan kapasitas minimal 1
	return &DynamicArray{
		Array:    make([]int, initCap),
		Capacity: initCap,
		Size:     0,
	}
}

func (da *DynamicArray) DoubleSize() {
	newCap := da.Capacity * 2
	newArray := make([]int, newCap)
	copy(newArray, da.Array[:da.Size]) // Salin hanya elemen valid
	da.Array = newArray
	da.Capacity = newCap
}

func (da *DynamicArray) Insert(value int) {
	if da.Size == da.Capacity {
		da.DoubleSize()
	}
	da.Array[da.Size] = value
	da.Size++
}

func (da *DynamicArray) Find(target int) int {
	for i := range da.Size { // Cari hanya sampai Size
		if da.Array[i] == target {
			return i
		}
	}
	return -1
}

func (da *DynamicArray) HalveSize() {
	newCap := max(da.Capacity/2, 1)
	newArray := make([]int, newCap)
	copy(newArray, da.Array[:da.Size]) // Salin hanya elemen valid
	da.Array = newArray
	da.Capacity = newCap
}

func (da *DynamicArray) Delete(target int) error {
	index := da.Find(target)
	if index == -1 {
		return fmt.Errorf("delete: %d not found", target)
	}

	// Geser elemen ke kiri untuk hapus target
	copy(da.Array[index:], da.Array[index+1:da.Size])
	da.Size--

	// Periksa apakah perlu mengecilkan kapasitas
	if da.Capacity > 1 && da.Size <= da.Capacity/4 {
		da.HalveSize()
	}
	return nil
}

func main() {
	da := NewDynamicArray(2)

	da.Insert(10)
	da.Insert(20)
	da.Insert(30)
	da.Insert(40)

	fmt.Println("Array:", da.Array[:da.Size]) // Output: [10 20 30 40]

	err := da.Delete(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After delete:", da.Array[:da.Size]) // Output: [10 30 40]
}
