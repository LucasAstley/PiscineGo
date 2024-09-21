package piscine

func IsLetter(s string) bool {
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
