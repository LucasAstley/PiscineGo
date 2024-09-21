package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 57; i >= 48; i-- {
		for j := 57; j >= 48; j-- {
			for k := 57; k >= 48; k-- {
				for l := 57; l >= 48; l-- {
					if k < i {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(32)
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if !(i == 48 && j == 49 && k == 48 && l == 48) {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					} else if i == k && l < j {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(32)
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if !(i == 48 && j == 49 && k == 48 && l == 48) {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	}
}
