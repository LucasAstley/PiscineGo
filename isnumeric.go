package piscine

func IsNumeric(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 48 || rune(array[0][i]) > 57 {
			return false
		}
	}
	return true
}
