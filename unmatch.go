package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		how_many := 1
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				how_many += 1
			}
		}
		if how_many%2 != 0 {
			return a[i]
		}
	}
	return -1
}
