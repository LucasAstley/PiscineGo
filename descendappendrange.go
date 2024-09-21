package piscine

func DescendAppendRange(max, min int) []int {
	array := []int{}
	if min == max || min > max {
		return array
	} else {
		for i := max; i > min; i-- {
			array = append(array, i)
		}
		return array
	}
}
