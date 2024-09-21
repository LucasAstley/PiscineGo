package piscine

func ActiveBits(n int) int {
	binary := decimalToBase(n, "01")
	counter := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			counter++
		}
	}
	return counter
}
