package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		str := string(stdin)
		if len(str) > 0 {
			for len(str) > 0 {
				z01.PrintRune(rune(str[0]))
				str = str[1:]
			}
		}
	} else {
		for i := 0; i < len(args); i++ {
			file, err := os.ReadFile(args[i])
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
}
