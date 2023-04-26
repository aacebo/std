package map_set_test

import (
	"testing"

	"github.com/aacebo/std/map_set"
)

func TestMap(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		t.Run("should have correct size", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)
			m.Set("b", 1)
			m.Set("c", 2)

			if m.Size() != 3 {
				t.Errorf("should have size 3")
			}
		})
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("should be empty", func(t *testing.T) {
			m := map_set.New[string, int]()

			if !m.Empty() {
				t.Errorf("should be empty")
			}
		})

		t.Run("should not be empty", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)

			if m.Empty() {
				t.Errorf("should not be empty")
			}
		})
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("should have elements", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)
			m.Set("b", 1)
			m.Set("c", 2)

			if !m.Has("a", "b", "c") {
				t.Errorf("should have a, b, c")
			}
		})

		t.Run("should not have element", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)
			m.Set("b", 1)
			m.Set("c", 2)

			if m.Has("d") {
				t.Errorf("should not have d")
			}
		})
	})

	t.Run("Set", func(t *testing.T) {
		t.Run("should set key/value", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)

			if m.Get("a") != 0 {
				t.Errorf("key a should have value 0")
			}

			m.Set("a", 10)

			if m.Get("a") != 10 {
				t.Errorf("key a should have value 10")
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete element", func(t *testing.T) {
			m := map_set.New[string, int]()
			m.Set("a", 0)

			if !m.Has("a") {
				t.Errorf("should have element a")
			}

			m.Delete("a")

			if m.Has("a") {
				t.Errorf("should not have element a")
			}
		})
	})
}
