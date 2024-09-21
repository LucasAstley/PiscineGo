package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		if str == "" {
			return "\n"
		}
		return "Invalid Output\n"
	}
	newStr := ""
	counter := 0
	for i, j := range str {
		if i == len(str)-1 && counter == 5 {
			newStr += " "
		} else if counter == 5 {
			newStr += " "
			counter = 0
		} else if j != ' ' {
			newStr += string(j)
			counter++
		} else {
			newStr += ""
		}
	}
	if len(newStr) != 0 {
		for newStr[len(newStr)-1] == ' ' {
			newStr = newStr[:len(newStr)-1]
		}
	}
	newStr += "\n"
	return newStr
}
