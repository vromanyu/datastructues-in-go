package linked_list

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type LinkedList[T constraints.Ordered] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

type Node[T constraints.Ordered] struct {
	Value T
	Next  *Node[T]
}

func NewLinkedList[T constraints.Ordered](value T) *LinkedList[T] {
	node := NewNode(value)
	return &LinkedList[T]{
		Head: node,
		Tail: node,
		size: 1,
	}
}

func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}

func NewEmptyIntLinkedList[T constraints.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func (list *LinkedList[T]) Append(value T) {
	newNode := NewNode(value)
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}
	list.size++
}

func (list *LinkedList[T]) RemoveLast() *Node[T] {
	if list.size == 0 {
		return nil
	}

	pre, temp := list.Head, list.Head

	for temp.Next != nil {
		pre = temp
		temp = temp.Next
	}
	list.Tail = pre
	list.Tail.Next = nil
	list.size--
	if list.size == 0 {
		list.Head = nil
		list.Tail = nil
	}
	return temp

}

func (list *LinkedList[T]) Prepend(data T) {
	newNode := NewNode(data)
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head = newNode
	}
	list.size++
}

func (list *LinkedList[T]) RemoveFirst() *Node[T] {
	if list.Head == nil {
		return nil
	}
	target := list.Head
	list.Head = target.Next
	target.Next = nil
	list.size--
	if list.size == 0 {
		list.Tail = nil
	}
	return target
}

func (list *LinkedList[T]) Get(index int) *Node[T] {
	if index < 0 || index >= list.size {
		return nil
	}
	if index == list.size-1 {
		return list.Tail
	}
	if index == 0 {
		return list.Head
	}
	temp := list.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp
}

func (list *LinkedList[T]) Set(index int, value T) bool {
	temp := list.Get(index)
	if temp == nil {
		return false
	}
	temp.Value = value
	return true
}

func (list *LinkedList[T]) Insert(index int, value T) bool {
	if index < 0 || index > list.size {
		return false
	}
	if index == 0 {
		list.Prepend(value)
		return true
	}
	if index == list.size {
		list.Append(value)
		return true
	}
	temp := list.Get(index - 1)
	node := NewNode[T](value)
	node.Next = temp.Next
	temp.Next = node
	list.size++
	return true
}

func (list *LinkedList[T]) Print() {
	start := list.Head
	for start != nil {
		fmt.Println(start.Value)
		start = start.Next
	}
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (node *Node[T]) String() string {
	return fmt.Sprintf("value: %v, next: %v", node.Value, node.Next)
}
