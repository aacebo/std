package queue

type Queue[T any] []T

func New[T any](items ...T) Queue[T] {
	self := Queue[T]{}
	self = append(self, items...)
	return self
}

func (self Queue[T]) Size() int {
	return len(self)
}

func (self Queue[T]) Empty() bool {
	return self.Size() == 0
}

func (self Queue[T]) Top() T {
	return self[0]
}

func (self Queue[T]) Bottom() T {
	return self[self.Size()-1]
}

func (self *Queue[T]) Push(items ...T) {
	*self = append(*self, items...)
}

func (self *Queue[T]) Pop() T {
	v := self.Top()
	arr := *self
	*self = arr[1:]
	return v
}
