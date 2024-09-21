package piscine

func AppendRange(min, max int) []int {
	array := []int{}
	if min == max || min > max {
		return []int(nil)
	} else {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
		return array
	}
}
