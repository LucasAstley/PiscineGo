package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	} else {
		newStr := ""
		for i := 2; i < len(str); i += 3 {
			newStr += string(rune(str[i]))
		}
		return newStr + "\n"
	}
}
