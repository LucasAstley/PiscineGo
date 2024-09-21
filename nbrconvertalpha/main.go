package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, i := range os.Args[1:] {
		if os.Args[1] == "--upper" {
			if ConvStrInt(i) > 0 && ConvStrInt(i) < 27 {
				z01.PrintRune(rune(ConvStrInt(i) + 64))
			}
		} else if ConvStrInt(i) > 0 && ConvStrInt(i) < 27 {
			z01.PrintRune(rune(ConvStrInt(i) + 96))
		} else {
			z01.PrintRune(rune(32))
		}
	}
	z01.PrintRune('\n')
}

func ConvStrInt(s string) int {
	allCaracter := []rune(s)
	result := 0
	for i := 0; i < len(allCaracter); i++ {
		allCaracter[i] = allCaracter[i] - 48
		diff := 1
		for j := 0; j < (len(allCaracter) - 1 - i); j++ {
			diff = 10
		}
		result += int(allCaracter[i]) * diff
	}
	return result
}
