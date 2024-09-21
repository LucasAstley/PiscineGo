package piscine

func Compact(ptr *[]string) int {
	counter := 0
	array := []string{}
	for _, i := range *ptr {
		if i != "" {
			counter++
			array = append(array, i)
		}
	}
	*ptr = array
	return counter
}
