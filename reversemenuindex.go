package piscine

func ReverseMenuIndex(menu []string) []string {
	newMenu := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		newMenu[len(newMenu)-1-i] = menu[i]
	}
	return newMenu
}
