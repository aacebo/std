package std

import "log"

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (self Map[K, V]) Size() int {
	return len(self)
}

func (self Map[K, V]) Empty() bool {
	return self.Size() == 0
}

func (self Map[K, V]) Has(keys ...K) bool {
	for _, k := range keys {
		if _, ok := self[k]; !ok {
			return false
		}
	}

	return true
}

func (self Map[K, V]) Get(key K) V {
	v, ok := self[key]

	if !ok {
		log.Panicf("key %v outside of range", key)
	}

	return v
}

func (self Map[K, V]) Set(key K, value V) {
	self[key] = value
}

func (self Map[K, V]) Delete(key K) {
	delete(self, key)
}
