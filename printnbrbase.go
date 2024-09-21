package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) >= 2 && IsValidBase(base) && dups(base) == false {
		array := []int{}
		if base == "0123456789" { // Numeric base
			PrintNbr(nbr)
		} else if base == "01" { // Binary base
			if nbr < 0 {
				z01.PrintRune('-')
				nbr *= -1
			}
			if nbr == 0 {
				z01.PrintRune('0')
			}
			for nbr > 0 {
				array = append(array, nbr%2)
				nbr /= 2
			}
			for i := len(array) - 1; i >= 0; i-- {
				z01.PrintRune(rune(array[i] + 48))
			}
		} else if base == "0123456789ABCDEF" { // Hexadecimal base
			if nbr < 0 {
				z01.PrintRune('-')
				nbr *= -1
			}
			if nbr == 0 {
				z01.PrintRune('0')
			}
			for nbr > 0 {
				array = append(array, nbr%16)
				nbr /= 16
			}
			for i := len(array) - 1; i >= 0; i-- {
				if array[i] >= 10 {
					z01.PrintRune(rune(array[i] + 55))
				} else {
					z01.PrintRune(rune(array[i] + 48))
				}
			}
		} else if IsValidBase(base) { // Custom base
			if nbr < 0 {
				z01.PrintRune('-')
				nbr *= -1
			}
			if nbr == 0 {
				z01.PrintRune(rune(base[0]))
			}
			for nbr > 0 {
				array = append(array, nbr%len(base))
				nbr /= len(base)
			}
			for i := len(array) - 1; i >= 0; i-- {
				z01.PrintRune(rune(base[array[i]]))
			}
		}
	} else {
		z01.PrintRune(78)
		z01.PrintRune(86)
	}
}

func IsValidBase(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 32 || rune(array[0][i]) > 126 || rune(array[0][i]) == 43 || rune(array[0][i]) == 45 {
			return false
		}
	}
	return true
}
