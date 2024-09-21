package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	puissance := 1
	digit_count := 0
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n *= -1
		}
	}
	for puissance <= n/10 {
		digit_count += 1
		puissance *= 10
	}
	for digit_count >= 0 {
		z01.PrintRune(rune((n / puissance) + 48))
		n %= puissance
		puissance /= 10
		digit_count -= 1
	}
}
