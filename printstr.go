package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for len(s) > 0 {
		z01.PrintRune(rune(s[0]))
		s = s[1:]
	}
}
