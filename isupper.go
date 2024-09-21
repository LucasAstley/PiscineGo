package piscine

func IsUpper(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 65 || rune(array[0][i]) > 90 {
			return false
		}
	}
	return true
}
