package piscine

func StringToIntSlice(str string) []int {
	var arr []int
	for _, char := range str {
		arr = append(arr, int(char))
	}
	return arr
}
