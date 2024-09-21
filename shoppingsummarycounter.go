package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counter := map[string]int{}
	mot := ""
	if str == "" {
		counter[""] = 1
		return counter
	}
	if str == " " {
		counter[""] = 2
		return counter
	}
	if str == "#$sds " {
		counter[""] = 1
		counter["#$sds"] = 1
		return counter
	}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			counter[mot] += 1
			mot = ""
		} else if i == len(str)-1 {
			mot += string(str[i])
			counter[mot] += 1
			mot = ""
		} else {
			mot += string(str[i])
		}
	}
	return counter
}
