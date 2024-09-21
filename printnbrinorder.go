package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	array := []int{}
	for n > 0 {
		lastnumber := n % 10
		n /= 10
		array = append(array, lastnumber)
	}
	SortIntegerTable(array)
	for i := 0; i < len(array); i++ {
		z01.PrintRune(rune(array[i] + 48))
	}
}
