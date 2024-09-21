package piscine

func Somme(index int) int {
	if index == 0 {
		return 0
	} else {
		return index + Somme(index-1)
	}
}
