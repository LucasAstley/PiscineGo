package piscine

func Split(s, sep string) []string {
	str := ""
	array := []string{}
	sepLen := len(sep)
	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			array = append(array, str)
			str = ""
			i += sepLen - 1
		} else {
			str += string(s[i])
		}
	}
	array = append(array, str)
	return array
}
