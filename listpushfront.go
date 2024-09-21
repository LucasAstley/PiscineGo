package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		headBackup := l.Head
		l.Head = &NodeL{Data: data}
		l.Tail = headBackup
	} else {
		n := &NodeL{Data: data}
		n.Next = l.Head
		l.Head = n
	}
}
