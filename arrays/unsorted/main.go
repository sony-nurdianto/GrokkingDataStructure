package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/sony-nurdianto/GrokkingDataStructure/arrays/unsorted/data"
)

func main() {
	// Inisialisasi UnsortedArray dengan kapasitas 10
	ua := data.NewUnsortedArray(10)

	fmt.Println("Menambahkan elemen ke dalam UnsortedArray:")
	for range 8 {
		val := rand.Intn(100) + 1 // Angka antara 1-100
		err := ua.Insert(val)
		if err != nil {
			log.Println("Gagal menambahkan elemen:", err)
		} else {
			fmt.Printf("Menambahkan: %d\n", val)
		}
	}

	// Menampilkan isi array
	fmt.Println("\nIsi array setelah penambahan:")
	ua.Traverse(func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()

	// Mencari elemen dalam array
	searchVal := rand.Intn(100) + 1
	fmt.Printf("\nMencari elemen %d dalam array...\n", searchVal)
	result := ua.Find(searchVal)
	if result != -1 {
		fmt.Printf("Elemen %d ditemukan!\n", searchVal)
	} else {
		fmt.Printf("Elemen %d tidak ditemukan.\n", searchVal)
	}

	// Menghapus elemen di indeks tertentu (misalnya indeks ke-2)
	deleteIndex := 2
	fmt.Printf("\nMenghapus elemen pada indeks %d...\n", deleteIndex)
	err := ua.Delete(deleteIndex)
	if err != nil {
		log.Println("Gagal menghapus elemen:", err)
	} else {
		fmt.Println("Elemen berhasil dihapus.")
	}

	// Menampilkan isi array setelah penghapusan
	fmt.Println("\nIsi array setelah penghapusan:")
	ua.Traverse(func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()
}
