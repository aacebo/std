package std

type Queue[T any] struct {
	items []*T
}

func NewQueue[T any]() *Queue[T] {
	v := Queue[T]{items: []*T{}}

	return &v
}

func (self *Queue[T]) Push(v T) {
	self.items = append(self.items, &v)
}

func (self *Queue[T]) Pop() *T {
	item := self.items[0]
	self.items = self.items[1:]
	return item
}

func (self *Queue[T]) Length() int {
	return len(self.items)
}

func (self *Queue[T]) At(i int) *T {
	return self.items[i]
}
