package queue_test

import (
	"testing"

	"github.com/aacebo/std/queue"
)

func TestQueue(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		t.Run("should have correct size", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Size() != 3 {
				t.Errorf("should have size of 3")
			}
		})
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("should be empty", func(t *testing.T) {
			q := queue.New[int]()

			if !q.Empty() {
				t.Errorf("should be empty")
			}
		})

		t.Run("should not be empty", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Empty() {
				t.Errorf("should not be empty")
			}
		})
	})

	t.Run("Top", func(t *testing.T) {
		t.Run("should be first element", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Top() != 0 {
				t.Errorf("should have top of 0")
			}
		})
	})

	t.Run("Bottom", func(t *testing.T) {
		t.Run("should be last element", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Bottom() != 2 {
				t.Errorf("should have bottom of 2")
			}
		})
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("should add new element to back", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Bottom() != 2 {
				t.Errorf("should have bottom of 2")
			}

			q.Push(5)

			if q.Bottom() != 5 {
				t.Errorf("should have bottom of 5")
			}

			if q.Size() != 4 {
				t.Errorf("should have size of 4")
			}
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("should remove element from front", func(t *testing.T) {
			q := queue.New(0, 1, 2)

			if q.Top() != 0 {
				t.Errorf("should have top of 0")
			}

			if q.Pop() != 0 {
				t.Errorf("should have popped off 0")
			}

			if q.Size() != 2 {
				t.Errorf("should have size of 2")
			}
		})
	})
}
