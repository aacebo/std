package set

type Set[T comparable] map[T]bool

func New[T comparable](items ...T) Set[T] {
	self := Set[T]{}

	for _, v := range items {
		self[v] = true
	}

	return self
}

func (self Set[T]) Size() int {
	return len(self)
}

func (self Set[T]) Empty() bool {
	return self.Size() == 0
}

func (self Set[T]) Has(items ...T) bool {
	for _, v := range items {
		if _, ok := self[v]; !ok {
			return false
		}
	}

	return true
}

func (self Set[T]) Add(items ...T) {
	for _, v := range items {
		self[v] = true
	}
}

func (self Set[T]) Delete(items ...T) {
	for _, v := range items {
		if _, ok := self[v]; ok {
			delete(self, v)
		}
	}
}
