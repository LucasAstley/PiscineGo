package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		curentNode := l.Head
		for curentNode.Next != nil {
			curentNode = curentNode.Next
		}
		curentNode.Next = n
	}
}
