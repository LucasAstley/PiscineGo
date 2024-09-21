package piscine

func Join(strs []string, sep string) string {
	newStr := ""
	for i := 0; i < len(strs)-1; i++ {
		newStr += strs[i]
		newStr += sep
	}
	newStr += strs[len(strs)-1]
	return newStr
}
