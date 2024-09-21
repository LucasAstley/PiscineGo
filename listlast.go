package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		curentNode := l.Head
		for curentNode.Next != nil {
			curentNode = curentNode.Next
		}
		return curentNode.Data
	}
}
