package main

import (
	"fmt"
	"os"
)

func main() {
	params := paramsToArray()
	toInsert := ""
	inOrder := false
	str := ""
	if len(params) == 0 {
		printHelp()
	}
	for i := 0; i < len(params); i++ {
		if len(params[i]) > 0 {
			if params[i][:1] == "-" {
				if params[i][:2] == "-i" {
					toInsert = params[0][3:]
				} else if params[i][:2] == "-o" {
					inOrder = true
				} else if params[i][:2] == "-h" {
					printHelp()
				} else if params[i][:3] == "--i" {
					toInsert = params[0][9:]
				} else if params[i][:3] == "--o" {
					inOrder = true
				} else if params[i][:3] == "--h" {
					printHelp()
				} else {
					printHelp()
				}
			} else {
				str = params[i]
			}
		} else {
			str = params[i]
		}
	}
	if inOrder {
		str += toInsert
		for i := 0; i < len(str); i++ {
			for j := i + 1; j < len(str); j++ {
				if str[i] > str[j] {
					temp := []rune(str)
					temp[i], temp[j] = temp[j], temp[i]
					str = string(temp)
				}
			}
		}
		fmt.Println(str)
	} else if len(str) > 0 {
		fmt.Println(str + toInsert)
	}
}

func paramsToArray() []string {
	params := os.Args
	args := []string{}
	for i := 1; i < len(params); i++ {
		tempString := ""
		for _, w := range params[i] {
			tempString += string(w)
		}
		args = append(args, tempString)
	}
	return args
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}
