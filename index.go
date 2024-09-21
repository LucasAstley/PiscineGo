package piscine

func Index(s string, toFind string) int {
	array := []string{s}
	for i := 0; i <= len(array[0])-len(toFind); i++ {
		if string(array[0][i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}
