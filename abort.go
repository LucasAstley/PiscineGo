package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	SortIntegerTable(array)
	return array[2]
}
