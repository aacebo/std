package structures

type Queue[T any] struct {
	values []*T
}

func NewQueue[T any]() *Queue[T] {
	v := Queue[T]{values: []*T{}}

	return &v
}

func (self *Queue[T]) Push(v T) {
	self.values = append(self.values, &v)
}

func (self *Queue[T]) Pop() *T {
	item := self.values[0]
	self.values = self.values[1:]
	return item
}

func (self *Queue[T]) Size() int {
	return len(self.values)
}

func (self *Queue[T]) At(i int) *T {
	return self.values[i]
}
