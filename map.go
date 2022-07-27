package std

type Map[T any] struct {
	items map[string]*T
}

func NewMap[T any]() *Map[T] {
	v := Map[T]{items: map[string]*T{}}

	return &v
}

func (self *Map[T]) Get(key string) *T {
	if v, ok := self.items[key]; ok {
		return v
	}

	return nil
}

func (self *Map[T]) Has(key string) bool {
	return self.Get(key) != nil
}

func (self *Map[T]) Set(key string, value T) {
	self.items[key] = &value
}

func (self *Map[T]) Remove(key string) {
	if self.Has(key) {
		delete(self.items, key)
	}
}
