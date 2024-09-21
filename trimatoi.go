package piscine

func TrimAtoi(s string) int {
	num := 0
	neg := false
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
				num *= 10
				num += int(rune(s[i])) - 48
			} else if rune(s[i]) == 45 {
				if num == 0 {
					neg = true
				}
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
