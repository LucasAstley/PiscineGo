package piscine

func Atoi(s string) int {
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
