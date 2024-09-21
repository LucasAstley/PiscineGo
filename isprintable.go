package piscine

func IsPrintable(s string) bool {
	array := []string{s}
	for i := 0; i <= len(array[0])-1; i++ {
		if rune(array[0][i]) < 32 || rune(array[0][i]) > 126 {
			return false
		}
	}
	return true
}
