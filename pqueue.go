package std

type PQueue[T any] struct {
	items   []T
	compare func(T, T) bool
}

func NewPQueue[T any](compare func(T, T) bool) PQueue[T] {
	self := PQueue[T]{[]T{}, compare}
	return self
}

func (self PQueue[T]) Size() int {
	return len(self.items)
}

func (self PQueue[T]) Empty() bool {
	return self.Size() == 0
}

func (self PQueue[T]) Top() T {
	return self.items[0]
}

func (self PQueue[T]) Bottom() T {
	return self.items[self.Size()-1]
}

func (self *PQueue[T]) Push(items ...T) {
	for _, v := range items {
		self.items = append(self.items, v)
		self.siftUp()
	}
}

func (self *PQueue[T]) Pop() T {
	v := self.Top()
	self.items = self.items[1:]
	return v
}

func (self *PQueue[T]) siftUp() {
	i := self.Size() - 1

	for i > 0 && self.compare(self.items[i], self.items[i-1]) {
		tmp := self.items[i]
		self.items[i] = self.items[i-1]
		self.items[i-1] = tmp
		i--
	}
}
