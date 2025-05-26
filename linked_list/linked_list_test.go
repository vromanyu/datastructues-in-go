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
