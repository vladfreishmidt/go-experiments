package genericsingostructures

// Implement a linked list that works with generics
type Node[T any] struct {
	Data T
	Next *Node[T]
}

type List[T any] struct {
	Start *Node[T]
}

func (l *List[T]) Add(data T) {
	n := Node[T]{
		Data: data,
		Next: nil,
	}

	if l.Start == nil {
		l.Start = &n
		return
	}

	if l.Start.Next == nil {
		l.Start.Next = &n
		return
	}

	temp := l.Start
	l.Start = l.Start.Next
	l.Add(data)
	l.Start = temp
}
