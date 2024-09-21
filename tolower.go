package piscine

func ToLower(s string) string {
	array := []string{s}
	new_word := ""
	for i := 0; i <= len(array[0])-1; i++ {
		if IsUpper(string(array[0][i])) {
			new_word += string(rune(array[0][i] + 32))
		} else {
			new_word += string(array[0][i])
		}
	}
	return new_word
}
