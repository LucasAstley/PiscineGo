package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i <= 57; i++ {
		for j := 48; j <= 57; j++ {
			for k := 48; k <= 57; k++ {
				for l := 48; l <= 57; l++ {
					if i < k {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(32)
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i == 57 && j == 56 && k == 57 && l == 57 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					} else if i == k && j < l {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(32)
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i == 57 && j == 56 && k == 57 && l == 57 {
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
}
