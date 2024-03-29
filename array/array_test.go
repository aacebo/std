package array_test

import (
	"testing"

	"github.com/aacebo/std/array"
)

func TestArray(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		t.Run("should have correct size", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}
		})
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("should be empty", func(t *testing.T) {
			v := array.New[int]()

			if !v.Empty() {
				t.Errorf("should be empty")
			}
		})

		t.Run("should not be empty", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Empty() {
				t.Errorf("should not be empty")
			}
		})
	})

	t.Run("First", func(t *testing.T) {
		t.Run("should have correct first element", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.First() != 0 {
				t.Errorf("first element %d should be %d", v.First(), 0)
			}
		})
	})

	t.Run("Last", func(t *testing.T) {
		t.Run("should have correct last element", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Last() != 2 {
				t.Errorf("last element %d should be %d", v.Last(), 2)
			}
		})
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("should push element to end", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v.Push(10)

			if v.Size() != 4 {
				t.Errorf("size %d should be %d", v.Size(), 4)
			}

			if v.Last() != 10 {
				t.Errorf("last element %d should be %d", v.Last(), 10)
			}
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("should pop element from the end", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			if e := v.Pop(); e != 2 {
				t.Errorf("element %d should be %d", e, 2)
			}

			if v.Size() != 2 {
				t.Errorf("size %d should be %d", v.Size(), 2)
			}
		})
	})

	t.Run("UnShift", func(t *testing.T) {
		t.Run("should unshift element to start", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v.UnShift(-1)

			if v.Size() != 4 {
				t.Errorf("size %d should be %d", v.Size(), 4)
			}

			if v.First() != -1 {
				t.Errorf("first element %d should be %d", v.First(), -1)
			}
		})
	})

	t.Run("Shift", func(t *testing.T) {
		t.Run("should shift element from the start", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			if e := v.Shift(); e != 0 {
				t.Errorf("element %d should be %d", e, 0)
			}

			if v.Size() != 2 {
				t.Errorf("size %d should be %d", v.Size(), 2)
			}
		})
	})

	t.Run("Splice", func(t *testing.T) {
		t.Run("should remove element", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v.Splice(1, 1)

			if v.Size() != 2 {
				t.Errorf("size %d should be %d", v.Size(), 2)
			}

			if v.First() != 0 || v.Last() != 2 {
				t.Errorf("elements should be 0, 2")
			}
		})

		t.Run("should add element", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v.Splice(1, 0, 25)

			if v.Size() != 4 {
				t.Errorf("size %d should be %d", v.Size(), 4)
			}

			if v.First() != 0 || v.Last() != 2 {
				t.Errorf("elements should be 0, 2")
			}

			if v[1] != 25 {
				t.Errorf("element %d should be %d", v[2], 25)
			}

			if v[2] != 1 {
				t.Errorf("element %d should be %d", v[2], 1)
			}
		})

		t.Run("should add element and remove element", func(t *testing.T) {
			v := array.New(0, 1, 2)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v.Splice(1, 1, 25, 50)

			if v.Size() != 4 {
				t.Errorf("size %d should be %d", v.Size(), 4)
			}

			if v.First() != 0 || v.Last() != 2 {
				t.Errorf("elements should be 0, 2")
			}

			if v[1] != 25 {
				t.Errorf("element %d should be %d", v[2], 25)
			}

			if v[2] != 50 {
				t.Errorf("element %d should be %d", v[2], 50)
			}
		})
	})

	t.Run("Sort", func(t *testing.T) {
		t.Run("should sort ascending", func(t *testing.T) {
			v := array.New(17, 1, 50)

			v.Sort(func(a int, b int) bool {
				return a < b
			})

			if v.First() != 1 || v.Last() != 50 {
				t.Errorf("elements should be in acending order")
			}
		})

		t.Run("should sort descending", func(t *testing.T) {
			v := array.New(17, 1, 50)

			v.Sort(func(a int, b int) bool {
				return a > b
			})

			if v.First() != 50 || v.Last() != 1 {
				t.Errorf("elements should be in descending order")
			}
		})
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("should find element index", func(t *testing.T) {
			v := array.New(17, 1, 50)
			i := v.Find(func(a int) int {
				if a == 1 {
					return 0
				} else if a > 1 {
					return 1
				}

				return -1
			})

			if i != 1 {
				t.Errorf("element %d should be at index %d", 1, 1)
			}

			i = v.Find(func(a int) int {
				if a == 50 {
					return 0
				} else if a > 50 {
					return 1
				}

				return -1
			})

			if i != 2 {
				t.Errorf("element %d should be at index %d", 50, 2)
			}
		})

		t.Run("should not find element", func(t *testing.T) {
			v := array.New(0, 1, 2)
			i := v.Find(func(a int) int {
				if a == 50 {
					return 0
				} else if a > 50 {
					return 1
				}

				return -1
			})

			if i != -1 {
				t.Errorf("element %d should not exist", 50)
			}
		})
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("should filter elements", func(t *testing.T) {
			v := array.New(17, 1, 50)

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			v = v.Filter(func(a int) bool {
				return a > 1
			})

			if v.Size() != 2 {
				t.Errorf("size %d should be %d", v.Size(), 2)
			}

			if v.First() != 17 || v.Last() != 50 {
				t.Errorf("elements should be 17, 50")
			}
		})
	})

	t.Run("Slice", func(t *testing.T) {
		t.Run("should get sub array", func(t *testing.T) {
			v := array.New(50, 1, 79, 1000)

			if v.Size() != 4 {
				t.Errorf("size %d should be %d", v.Size(), 4)
			}

			v = v.Slice(1, v.Size())

			if v.Size() != 3 {
				t.Errorf("size %d should be %d", v.Size(), 3)
			}

			i := v.Find(func(a int) int {
				if a == 50 {
					return 0
				} else if a > 50 {
					return 1
				}

				return -1
			})

			if i > -1 {
				t.Errorf("element %d should not exist at index %d", 50, i)
			}
		})
	})
}
