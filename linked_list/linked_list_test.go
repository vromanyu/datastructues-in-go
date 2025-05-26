package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedListEmptyConstructor(t *testing.T) {
	list := NewEmptyIntLinkedList[int]()
	if list.Head != nil || list.Tail != nil || list.size != 0 {
		t.Errorf("expected empty linked list, got head: %v, tail: %v, size: %d", list.Head, list.Tail, list.size)
	}
}

func TestLinkedListConstructorWithOneValue(t *testing.T) {
	testCases := []struct {
		value int
	}{
		{value: 1},
		{value: 42},
		{value: 10},
		{value: 0},
		{value: 1000},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("create linked list with single value: %v\n", tc.value), func(t *testing.T) {
			list := NewLinkedList(tc.value)
			if list.Head == nil || list.Tail == nil || list.size != 1 {
				t.Errorf("expected linked list with one node, got head: %v, tail: %v, size: %d", list.Head, list.Tail, list.size)
			}
			if list.Head.Value != tc.value || list.Tail.Value != tc.value {
				t.Errorf("expected head and tail value to be %v, got head: %v, tail: %v", tc.value, list.Head.Value, list.Tail.Value)
			}
		})
	}
}

func TestLinkedListAppendToNotEmptyList(t *testing.T) {
	ll := NewLinkedList(1)
	ll.Append(2)
	if ll.size != 2 {
		t.Errorf("expected linked list size to be 2 after appending, got size: %d", ll.size)
	}
	if ll.Head.Next == nil {
		t.Errorf("expected head to have a next node, got nil")
	}
	if ll.Tail.Value != 2 {
		t.Errorf("expected tail value to be 2 after appending, got tail: %v", ll.Tail.Value)
	}
}

func TestLinkedListAppendToEmptyList(t *testing.T) {
	ll := NewEmptyIntLinkedList[int]()
	ll.Append(1)
	if ll.Head == nil || ll.Tail == nil || ll.size != 1 {
		t.Errorf("expected linked list with one node, got head: %v, tail: %v, size: %d", ll.Head, ll.Tail, ll.size)
	}
}

func TestRemoveLastFromEmptyList(t *testing.T) {
	ll := NewEmptyIntLinkedList[int]()
	node := ll.RemoveLast()
	if node != nil {
		t.Errorf("expected nil when removing last from empty list, got: %v", node)
	}
}

func TestRemoveLastFromListWithOneNode(t *testing.T) {
	ll := NewLinkedList(1)
	node := ll.RemoveLast()
	if node == nil || node.Value != 1 {
		t.Errorf("expected to remove node with value 1, got: %v", node)
	}
	if ll.size != 0 || ll.Head != nil || ll.Tail != nil {
		t.Errorf("expected linked list to be empty after removing last node, got size: %d, head: %v, tail: %v", ll.size, ll.Head, ll.Tail)
	}
}

func TestRemoveLastFromListWithMultipleNodes(t *testing.T) {
	ll := NewLinkedList(1)
	ll.Append(2)
	ll.Append(3)

	node := ll.RemoveLast()
	if node == nil || node.Value != 3 {
		t.Errorf("expected to remove node with value 3, got: %v", node)
	}
	if ll.size != 2 || ll.Tail.Value != 2 {
		t.Errorf("expected linked list size to be 2 and tail value to be 2, got size: %d, tail: %v", ll.size, ll.Tail.Value)
	}

	node = ll.RemoveLast()
	if node == nil || node.Value != 2 {
		t.Errorf("expected to remove node with value 2, got: %v", node)
	}
	if ll.size != 1 || ll.Tail.Value != 1 {
		t.Errorf("expected linked list size to be 1 and tail value to be 1, got size: %d, tail: %v", ll.size, ll.Tail.Value)
	}
	node = ll.RemoveLast()
	if node == nil || node.Value != 1 {
		t.Errorf("expected to remove node with value 1, got: %v", node)
	}
	if ll.size != 0 || ll.Head != nil || ll.Tail != nil {
		t.Errorf("expected linked list to be empty after removing last node, got size: %d, head: %v, tail: %v", ll.size, ll.Head, ll.Tail)
	}

	if ll.RemoveLast() != nil {
		t.Errorf("expected nil when removing last from empty list, got non-nil value")
	}
}
