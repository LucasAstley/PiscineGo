package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	sortParams(params)
	for i := 1; i < len(params); i++ {
		for _, w := range params[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func sortParams(table []string) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		min_arr := i
		for j := i + 1; j < n; j++ {
			if table[j] < table[min_arr] {
				min_arr = j
			}
		}
		table[min_arr], table[i] = table[i], table[min_arr]
	}
}
