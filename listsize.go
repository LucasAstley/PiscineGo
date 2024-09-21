package piscine

func ListSize(l *List) int {
	counter := 0
	n := l.Head
	for n != nil {
		counter++
		n = n.Next
	}
	return counter
}
