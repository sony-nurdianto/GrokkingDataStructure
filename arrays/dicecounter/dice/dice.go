package dice

import "fmt"

// DiceCounter menyimpan jumlah kemunculan angka dadu (1-6)
type DiceCounter struct {
	Counts [6]int
}

// RecordRoll mencatat hasil lemparan dadu
func (dc *DiceCounter) RecordRoll(val int) {
	if val < 1 || val > 6 {
		fmt.Println("Invalid roll:", val)
		return
	}
	dc.Counts[val-1]++
}

// MaxInArray mencari angka dadu yang paling sering muncul
func (dc *DiceCounter) MaxInArray() (int, int) {
	maxIndex, maxValue := 0, dc.Counts[0]
	for i, v := range dc.Counts {
		if v > maxValue {
			maxIndex, maxValue = i, v
		}
	}
	return maxIndex + 1, maxValue
}

// MinInArray mencari angka dadu yang paling jarang muncul
func (dc *DiceCounter) MinInArray() (int, int) {
	minIndex, minValue := 0, dc.Counts[0]
	for i, v := range dc.Counts {
		if v < minValue {
			minIndex, minValue = i, v
		}
	}
	return minIndex + 1, minValue
}
