package piscine

func Max(a []int) int {
	maxi := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxi] {
			maxi = i
		}
	}
	return a[maxi]
}
