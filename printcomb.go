package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i <= 57; i++ {
		for j := 49; j <= 57; j++ {
			for k := 50; k <= 57; k++ {
				if i < j && j < k {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					if i == 55 && j == 56 && k == 57 {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}
