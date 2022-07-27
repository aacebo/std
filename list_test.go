package std

import "testing"

func TestNewList(t *testing.T) {
	list := NewList[int]()

	if list.Size() != 0 {
		t.Fatal("list should initialize to empty")
	}
}
