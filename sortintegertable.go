package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		min_arr := i
		for j := i + 1; j < n; j++ {
			if table[j] < table[min_arr] {
				min_arr = j
			}
		}
		table[min_arr], table[i] = table[i], table[min_arr]
	}
}
