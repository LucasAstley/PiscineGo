package piscine

func Capitalize(s string) string {
	str := ""
	newWord := true
	for i := 0; i <= len(s)-1; i++ {
		if newWord {
			if isLetter(string(s[i])) {
				if i >= 1 {
					if isNumeric(string(s[i-1])) {
						str += toLower(string(s[i]))
						newWord = false
					} else {
						str += toUpper(string(s[i]))
						newWord = false
					}
				} else {
					str += toUpper(string(s[i]))
					newWord = false
				}
			} else {
				str += string(s[i])
			}
		} else if isLetter(string(s[i])) == false {
			str += string(s[i])
			newWord = true
		} else if isUpper(string(s[i])) {
			str += toLower(string(s[i]))
		} else {
			str += string(s[i])
		}
	}
	return str
}

func isLetter(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 65 {
			return false
		} else if rune(array[0][i]) > 90 && rune(array[0][i]) < 97 {
			return false
		} else if rune(array[0][i]) > 122 {
			return false
		}
	}
	return true
}

func toUpper(s string) string {
	array := []string{s}
	new_word := ""
	for i := 0; i <= len(array[0])-1; i++ {
		if isLower(string(array[0][i])) {
			new_word += string(rune(array[0][i] - 32))
		} else {
			new_word += string(array[0][i])
		}
	}
	return new_word
}

func isLower(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 97 || rune(array[0][i]) > 122 {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 65 || rune(array[0][i]) > 90 {
			return false
		}
	}
	return true
}

func toLower(s string) string {
	array := []string{s}
	new_word := ""
	for i := 0; i <= len(array[0])-1; i++ {
		if isUpper(string(array[0][i])) {
			new_word += string(rune(array[0][i] + 32))
		} else {
			new_word += string(array[0][i])
		}
	}
	return new_word
}

func isNumeric(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 48 || rune(array[0][i]) > 57 {
			return false
		}
	}
	return true
}
