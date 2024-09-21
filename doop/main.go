package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	operator := ""
	if len(args) == 3 {
		if atoi(args[0]) > 9223372036854775800 || atoi(args[2]) > 9223372036854775800 || atoi(args[0]) < -9223372036854775800 || atoi(args[2]) < -9223372036854775800 {
			return
		}
		for i := 0; i < len(args[0]); i++ {
			if (rune(args[0][i]) < 48 || rune(args[0][i]) > 57) && args[0][i] != '+' && args[0][i] != '-' {
				return
			}
		}
		for i := 0; i < len(args[2]); i++ {
			if (rune(args[0][i]) < 48 || rune(args[0][i]) > 57) && args[0][i] != '+' && args[0][i] != '-' {
				return
			}
		}
		if args[1] == "+" || args[1] == "-" || args[1] == "/" || args[1] == "*" || args[1] == "%" {
			operator = args[1]
		}
		if operator == "+" {
			os.Stdout.Write(intToByte(atoi(args[0]) + atoi(args[2])))
			os.Stdout.Write([]byte("\n"))
		} else if operator == "-" {
			os.Stdout.Write(intToByte(atoi(args[0]) - atoi(args[2])))
			os.Stdout.Write([]byte("\n"))
		} else if operator == "/" {
			if atoi(args[2]) == 0 {
				os.Stdout.Write([]byte("No division by 0"))
				os.Stdout.Write([]byte("\n"))
				return
			} else {
				os.Stdout.Write(intToByte(atoi(args[0]) / atoi(args[2])))
				os.Stdout.Write([]byte("\n"))
			}
		} else if operator == "*" {
			os.Stdout.Write(intToByte(atoi(args[0]) * atoi(args[2])))
			os.Stdout.Write([]byte("\n"))
		} else if operator == "%" {
			if atoi(args[2]) == 0 {
				os.Stdout.Write([]byte("No modulo by 0"))
				os.Stdout.Write([]byte("\n"))
				return
			} else {
				os.Stdout.Write(intToByte(atoi(args[0]) % atoi(args[2])))
				os.Stdout.Write([]byte("\n"))
			}
		}
	}
}

func atoi(s string) int {
	num := 0
	neg := false
	if len(s) > 0 {
		if s[0] == '+' {
			s = s[1:]
		} else if s[0] == '-' {
			neg = true
			s = s[1:]
		}
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
				num *= 10
				num += int(rune(s[i])) - 48
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	if neg {
		return -num
	} else {
		return num
	}
}

func intToByte(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var b []byte
	for n > 0 {
		b = append(b, byte(n%10)+'0')
		n /= 10
	}
	if neg {
		b = append(b, '-')
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
