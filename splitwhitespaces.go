package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	array := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " || string(s[i]) == "\n" || string(s[i]) == "	" {
			if string(s[i-1]) != " " {
				array = append(array, str)
				str = ""
			}
		} else {
			str += string(s[i])
		}
	}
	if string(str) != "" && string(str) != " " {
		array = append(array, str)
	}
	return array
}
