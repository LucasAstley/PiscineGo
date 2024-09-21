package piscine

func NRune(s string, n int) rune {
	array := []rune(s)
	if n > len(array) || n <= 0 {
		return rune('\x00')
	} else {
		return array[n-1]
	}
}
