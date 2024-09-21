package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i][4:])
		file, err := os.ReadFile(args[i][3:])
		if err != nil {
			errMsg := "ERROR: " + err.Error()
			for len(errMsg) > 0 {
				z01.PrintRune(rune(errMsg[0]))
				errMsg = errMsg[1:]
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		for len(file) > 0 {
			z01.PrintRune(rune(file[0]))
			file = file[1:]
		}
	}
}
