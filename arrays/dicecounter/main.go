package main

import (
	"fmt"
	"math/rand"

	"github.com/sony-nurdianto/GrokkingDataStructure/arrays/dicecounter/dice"
)

func main() {
	// Buat counter dadu
	diceCounter := &dice.DiceCounter{}

	// Simulasi lemparan dadu sebanyak 1000 kali
	totalRolls := 1000
	for range totalRolls {
		roll := rand.Intn(6) + 1 // Hasil antara 1-6
		diceCounter.RecordRoll(roll)
	}

	// Tampilkan distribusi angka dadu
	fmt.Println("Distribusi angka dadu setelah", totalRolls, "lemparan:")
	for i, count := range diceCounter.Counts {
		fmt.Printf("Angka %d muncul %d kali\n", i+1, count)
	}

	// Temukan angka yang paling sering muncul
	maxNum, maxCount := diceCounter.MaxInArray()
	fmt.Printf("\nAngka yang paling sering muncul: %d (%d kali)\n", maxNum, maxCount)

	// Temukan angka yang paling jarang muncul
	minNum, minCount := diceCounter.MinInArray()
	fmt.Printf("Angka yang paling jarang muncul: %d (%d kali)\n", minNum, minCount)
}
