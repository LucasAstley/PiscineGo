package piscine

func LastRune(s string) rune {
	array := []rune(s)
	return array[len(array)-1]
}
