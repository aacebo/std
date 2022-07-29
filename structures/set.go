package structures

import "github.com/aacebo/std"

type Set[T string | std.Number] struct {
	values []*T
}

func NewSet[T string | std.Number]() *Set[T] {
	v := Set[T]{values: []*T{}}

	return &v
}

func (self *Set[T]) Add(v T) {
	for i := 0; i < len(self.values); i++ {
		if *self.values[i] == v {
			return
		}
	}

	self.values = append(self.values, &v)
}

func (self *Set[T]) Remove(i int) {
	self.values = append(self.values[:i], self.values[i+1:]...)
}

func (self *Set[T]) At(i int) *T {
	return self.values[i]
}

func (self *Set[T]) Size() int {
	return len(self.values)
}
