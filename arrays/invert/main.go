package main

import "fmt"

func Invert[T any](arr []T) []T {
	n := len(arr) - 1

	for i := range n / 2 {
		arr[i], arr[n-i] = arr[n-i], arr[i]
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr = Invert(arr)

	fmt.Println(arr)
}
