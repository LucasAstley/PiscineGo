package piscine

func BasicAtoi(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num *= 10
		num += int(rune(s[i])) - 48
	}
	return num
}
