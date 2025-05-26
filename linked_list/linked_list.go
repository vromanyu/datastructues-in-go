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
