package piscine

func AlphaCount(s string) int {
	array := []string{s}
	compteur := 0
	for i := 0; i <= len(array[0])-1; i++ {
		if (rune(array[0][i]) >= 97 && rune(array[0][i]) <= 122) || (rune(array[0][i]) >= 65 && rune(array[0][i]) <= 90) {
			compteur += 1
		}
	}
	return compteur
}
