package piscine

func BasicAtoi2(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return 0
		}
		if rune(s[i]) > 57 || rune(s[i]) < 48 {
			return 0
		} else {
			num *= 10
			num += int(rune(s[i])) - 48
		}
	}
	return num
}
