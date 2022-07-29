package structures

import "testing"

func TestNewList(t *testing.T) {
	list := NewList[int]()

	if list.Size() != 0 {
		t.Fatal("list should initialize to empty")
	}
}

func TestListPush(t *testing.T) {
	list := NewList[int]()
	list.Push(1)
	list.Push(2)

	if list.Size() != 2 {
		t.Fatal("list size should increment with items added")
	}
}

func TestListPop(t *testing.T) {
	list := NewList[int]()
	list.Push(1)
	list.Push(2)
	list.Pop()

	if list.Size() != 1 {
		t.Fatal("list size should decrement when items popped")
	}
}

func TestListAt(t *testing.T) {
	list := NewList[int]()
	list.Push(1)
	list.Push(2)
	list.Push(3)

	if list.At(0) != nil && *list.At(0) == 1 {
		t.Fatal("list index 0 should have value 1")
	}

	if list.At(1) != nil && *list.At(1) != 2 {
		t.Fatal("list index 1 should have value 2")
	}

	if list.At(2) != nil && *list.At(2) != 3 {
		t.Fatal("list index 2 should have value 3")
	}
}
