package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(file))
	} else if len(args) < 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}
