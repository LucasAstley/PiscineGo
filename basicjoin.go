package piscine

func BasicJoin(elems []string) string {
	newStr := ""
	for i := 0; i < len(elems); i++ {
		newStr += elems[i]
	}
	return newStr
}
