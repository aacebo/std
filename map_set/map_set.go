package map_set

import "log"

type MapSet[K comparable, V any] map[K]V

func New[K comparable, V any]() MapSet[K, V] {
	return MapSet[K, V]{}
}

func (self MapSet[K, V]) Size() int {
	return len(self)
}

func (self MapSet[K, V]) Empty() bool {
	return self.Size() == 0
}

func (self MapSet[K, V]) Has(keys ...K) bool {
	for _, k := range keys {
		if _, ok := self[k]; !ok {
			return false
		}
	}

	return true
}

func (self MapSet[K, V]) Get(key K) V {
	v, ok := self[key]

	if !ok {
		log.Panicf("key %v outside of range", key)
	}

	return v
}

func (self MapSet[K, V]) Set(key K, value V) {
	self[key] = value
}

func (self MapSet[K, V]) Delete(key K) {
	delete(self, key)
}
