package list_test

import (
	"testing"

	"github.com/aacebo/std/list"
)

func TestList(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		t.Run("should have correct size", func(t *testing.T) {
			l := list.New[int]()

			if l.Size() != 0 {
				t.Errorf("size %d should be %d", l.Size(), 0)
			}

			l.Push(0, 1, 2)

			if l.Size() != 3 {
				t.Errorf("size %d should be %d", l.Size(), 3)
			}
		})
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("should be empty", func(t *testing.T) {
			l := list.New[int]()

			if !l.Empty() {
				t.Errorf("should be empty")
			}
		})

		t.Run("should not be empty", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.Empty() {
				t.Errorf("should not be empty")
			}
		})
	})

	t.Run("First", func(t *testing.T) {
		t.Run("should have valid first element", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.First() != 0 {
				t.Errorf("first element %d should be %d", l.First(), 0)
			}
		})
	})

	t.Run("Last", func(t *testing.T) {
		t.Run("should have valid last element", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.Last() != 2 {
				t.Errorf("last element %d should be %d", l.Last(), 2)
			}
		})
	})

	t.Run("At", func(t *testing.T) {
		t.Run("should have valid element at index", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.At(0) != 0 {
				t.Errorf("element at %d with value %d should have value %d", 0, l.At(0), 0)
			}

			if l.At(1) != 1 {
				t.Errorf("element at %d with value %d should have value %d", 1, l.At(1), 1)
			}

			if l.At(2) != 2 {
				t.Errorf("element at %d with value %d should have value %d", 2, l.At(2), 2)
			}
		})
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("should push elements to back", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.Size() != 3 {
				t.Errorf("size %d should be %d", l.Size(), 3)
			}

			l.Push(50, 20, 30)

			if l.Size() != 6 {
				t.Errorf("size %d should be %d", l.Size(), 6)
			}

			if l.First() != 0 {
				t.Errorf("first element %d should be %d", l.First(), 0)
			}

			if l.Last() != 30 {
				t.Errorf("last element %d should be %d", l.Last(), 30)
			}
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("should pop element from back", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.Size() != 3 {
				t.Errorf("size %d should be %d", l.Size(), 3)
			}

			if e := l.Pop(); e != 2 {
				t.Errorf("popped element %d should be %d", e, 2)
			}

			if l.Size() != 2 {
				t.Errorf("size %d should be %d", l.Size(), 2)
			}

			if l.First() != 0 {
				t.Errorf("first element %d should be %d", l.First(), 0)
			}

			if l.Last() != 1 {
				t.Errorf("last element %d should be %d", l.Last(), 1)
			}
		})
	})

	t.Run("UnShift", func(t *testing.T) {
		t.Run("should unshift elements to front", func(t *testing.T) {
			l := list.New[int]()
			l.Push(0, 1, 2)

			if l.Size() != 3 {
				t.Errorf("size %d should be %d", l.Size(), 3)
			}

			l.UnShift(50, 20, 30)

			if l.Size() != 6 {
				t.Errorf("size %d should be %d", l.Size(), 6)
			}

			if l.First() != 50 {
				t.Errorf("first element %d should be %d", l.First(), 50)
			}

			if l.Last() != 2 {
				t.Errorf("last element %d should be %d", l.Last(), 2)
			}
		})
	})

	t.Run("Shift", func(t *testing.T) {
		t.Run("should shift element from front", func(t *testing.T) {
			l := list.New[int]()
			l.UnShift(0, 1, 2)

			if l.Size() != 3 {
				t.Errorf("size %d should be %d", l.Size(), 3)
			}

			if e := l.Shift(); e != 0 {
				t.Errorf("popped element %d should be %d", e, 0)
			}

			if l.Size() != 2 {
				t.Errorf("size %d should be %d", l.Size(), 2)
			}

			if l.First() != 1 {
				t.Errorf("first element %d should be %d", l.First(), 1)
			}

			if l.Last() != 2 {
				t.Errorf("last element %d should be %d", l.Last(), 2)
			}
		})
	})

	t.Run("Splice", func(t *testing.T) {
		t.Run("should remove elements", func(t *testing.T) {
			l := list.New[int]()
			l.Push(50, 75, 100, 200)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			l.Splice(1, 2)

			if l.Size() != 2 {
				t.Errorf("size %d should be %d", l.Size(), 2)
			}

			if l.First() != 50 {
				t.Errorf("first element %d should be %d", l.First(), 50)
			}

			if l.Last() != 200 {
				t.Errorf("last element %d should be %d", l.Last(), 200)
			}
		})

		t.Run("should add elements", func(t *testing.T) {
			l := list.New[int]()
			l.Push(50, 75, 100, 200)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			l.Splice(1, 0, 1000, 2000)

			if l.Size() != 6 {
				t.Errorf("size %d should be %d", l.Size(), 6)
			}

			if l.First() != 50 {
				t.Errorf("first element %d should be %d", l.First(), 50)
			}

			if l.Last() != 200 {
				t.Errorf("last element %d should be %d", l.Last(), 200)
			}

			if l.At(1) != 1000 {
				t.Errorf("element at %d with value %d should be %d", 1, l.At(1), 1000)
			}

			if l.At(2) != 2000 {
				t.Errorf("element at %d with value %d should be %d", 2, l.At(2), 2000)
			}

			if l.At(3) != 75 {
				t.Errorf("element at %d with value %d should be %d", 3, l.At(3), 75)
			}
		})

		t.Run("should add and remove elements", func(t *testing.T) {
			l := list.New[int]()
			l.Push(50, 75, 100, 200)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			l.Splice(1, 1, 1000, 2000)

			if l.Size() != 5 {
				t.Errorf("size %d should be %d", l.Size(), 6)
			}

			if l.First() != 50 {
				t.Errorf("first element %d should be %d", l.First(), 50)
			}

			if l.Last() != 200 {
				t.Errorf("last element %d should be %d", l.Last(), 200)
			}

			if l.At(1) != 1000 {
				t.Errorf("element at %d with value %d should be %d", 1, l.At(1), 1000)
			}

			if l.At(2) != 2000 {
				t.Errorf("element at %d with value %d should be %d", 2, l.At(2), 2000)
			}

			if l.At(3) != 100 {
				t.Errorf("element at %d with value %d should be %d", 3, l.At(3), 75)
			}
		})

		t.Run("should remove head", func(t *testing.T) {
			l := list.New[int]()
			l.Push(50, 75, 100, 200)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			l.Splice(0, 1, 20)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			if l.First() != 20 {
				t.Errorf("first element %d should be %d", l.First(), 20)
			}
		})

		t.Run("should remove tail", func(t *testing.T) {
			l := list.New[int]()
			l.Push(50, 75, 100, 200)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			l.Splice(3, 1, 20)

			if l.Size() != 4 {
				t.Errorf("size %d should be %d", l.Size(), 4)
			}

			if l.Last() != 20 {
				t.Errorf("last element %d should be %d", l.Last(), 20)
			}
		})
	})

	t.Run("ForEach", func(t *testing.T) {
		t.Run("should iterate over all elements", func(t *testing.T) {
			i := 0
			l := list.New[int]()
			l.Push(50, 75, 1000)

			l.ForEach(func(j int, _ int) {
				if i != j {
					t.Errorf("index %d should be equal to %d", j, i)
				}

				i++
			})

			if i != 3 {
				t.Errorf("should have been called %d times", l.Size())
			}
		})
	})
}
