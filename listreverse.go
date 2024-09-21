package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	for i := 0; i < ListSize(l)/2; i++ {
		l.Head, l.Tail = l.Tail, l.Head
	}
}
