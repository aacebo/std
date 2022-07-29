package structures

type Map[T any] struct {
	values map[string]*T
}

func NewMap[T any]() *Map[T] {
	v := Map[T]{values: map[string]*T{}}

	return &v
}

func (self *Map[T]) Get(key string) *T {
	if v, ok := self.values[key]; ok {
		return v
	}

	return nil
}

func (self *Map[T]) Has(key string) bool {
	return self.Get(key) != nil
}

func (self *Map[T]) Set(key string, value T) {
	self.values[key] = &value
}

func (self *Map[T]) Remove(key string) {
	if self.Has(key) {
		delete(self.values, key)
	}
}

func (self *Map[T]) Size() int {
	return len(self.values)
}
