package linked_list

type LinkedList[T int | float64] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func NewLinkedList[T int | float64](value T) *LinkedList[T] {
	node := NewNode(value)
	return &LinkedList[T]{
		Head: node,
		Tail: node,
		size: 1,
	}
}

func NewEmptyIntLinkedList() *LinkedList[int] {
	return &LinkedList[int]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func NewEmptyFloatLinkedList() *LinkedList[float64] {
	return &LinkedList[float64]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

type Node[T int | float64] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T int | float64](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}
