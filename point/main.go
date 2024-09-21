package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	printStr("x = ")
	z01.PrintRune(52)
	z01.PrintRune(50)
	printStr(", y = ")
	z01.PrintRune(50)
	z01.PrintRune(49)
	printStr("\n")
}

func printStr(s string) {
	for len(s) > 0 {
		z01.PrintRune(rune(s[0]))
		s = s[1:]
	}
}
