package piscine

func StrRev(s string) string {
	mot := ""
	for i := len(s) - 1; i >= 0; i-- {
		mot += string(s[i])
	}
	return mot
}
