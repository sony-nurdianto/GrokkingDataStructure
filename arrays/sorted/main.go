package main

import (
	"fmt"

	"github.com/sony-nurdianto/GrokkingDataStructure/arrays/sorted/data"
)

func main() {
	sa := data.NewSortedArray(3)

	sa.Insert(8)
	sa.Insert(9)
	sa.Insert(9)

	fmt.Println(sa)

	sa.Delete(9)

	fmt.Println(sa)
}
