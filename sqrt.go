package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 0
	} else {
		for i := 0; i <= nb; i++ {
			if i > nb {
				return 0
			} else if i*i == nb {
				return i
			}
		}
	}
	return 0
}
