package set_test

import (
	"testing"

	"github.com/aacebo/std/set"
)

func TestSet(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		t.Run("should have correct size", func(t *testing.T) {
			s := set.New(1, 2, 3, 4, 5, 2)

			if s.Size() != 5 {
				t.Errorf("should have size %d", 5)
			}
		})
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("should be empty", func(t *testing.T) {
			s := set.New[int]()

			if !s.Empty() {
				t.Errorf("should be empty")
			}
		})

		t.Run("should not be empty", func(t *testing.T) {
			s := set.New(0, 1, 2)

			if s.Empty() {
				t.Error("should not be empty")
			}
		})
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("should have elements", func(t *testing.T) {
			s := set.New(0, 1, 2)

			if !s.Has(0, 1, 2) {
				t.Errorf("should have elements")
			}
		})

		t.Run("should not have element", func(t *testing.T) {
			s := set.New(0, 1, 2)

			if s.Has(0, 2, 7) {
				t.Errorf("should not have 7")
			}
		})
	})

	t.Run("Add", func(t *testing.T) {
		t.Run("should add new elements", func(t *testing.T) {
			s := set.New(0, 1, 2)
			s.Add(1, 2, 3)

			if !s.Has(0, 1, 2, 3) {
				t.Errorf("should have 0, 1, 2, 3")
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete elements", func(t *testing.T) {
			s := set.New(0, 1, 2)
			s.Delete(1, 2, 3)

			if s.Has(1, 2) {
				t.Errorf("should not have 1, 2")
			}
		})
	})
}
