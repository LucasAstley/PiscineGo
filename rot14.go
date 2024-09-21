package piscine

func Rot14(s string) string {
	new_string := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			if 122-rune(s[i]) < 14 {
				new_string += string(97 + (13 - (122 - rune(s[i]))))
			} else {
				new_string += string(rune(s[i]) + 14)
			}
		} else if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			if 90-rune(s[i]) < 14 {
				new_string += string(65 + (13 - (90 - rune(s[i]))))
			} else {
				new_string += string(rune(s[i]) + 14)
			}
		} else {
			new_string += string(s[i])
		}
	}
	return new_string
}
