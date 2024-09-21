package piscine

import "github.com/01-edu/z01"

func GenerateCombinations() {
	for i := 97; i <= 122; i++ { // Combination 1: Alphabet in Ascending Order
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
	for i := 65; i <= 90; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
	for i := 48; i <= 57; i++ { // Combination 2: Mixed Alphabet and Numbers
		z01.PrintRune(rune(i))
	}
	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
	for i := 48; i <= 57; i++ { // Combination 3: Numbers in Ascending and Descending Order
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
	for i := 57; i >= 48; i-- {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
	for i := 99; i <= 122; i++ { // Combination 4: Character Triplets
		z01.PrintRune(rune(i - 2))
		z01.PrintRune(rune(i - 1))
		z01.PrintRune(rune(i))
		if i == 122 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(44)
			z01.PrintRune(32)
		}

	}
}
