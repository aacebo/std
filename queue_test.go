package std

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int]()

	if queue.Size() != 0 {
		t.Fatal("queue must be empty by default")
	}
}

func TestQueuePush(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)

	if queue.Size() != 1 {
		t.Fatal("queue size must match push count")
	}
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	removed := queue.Pop()

	if queue.Size() != 1 {
		t.Fatal("queue size must match push count")
	}

	if *removed != 1 {
		t.Fatal("removed value should be the first added")
	}
}

func TestQueueAt(t *testing.T) {
	queue := NewQueue[string]()
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")

	if *queue.At(1) != "second" {
		t.Fatal("index value does not match expected")
	}
}
