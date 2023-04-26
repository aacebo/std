package array

import "sort"

type Array[T any] []T

func New[T any](items ...T) Array[T] {
	self := Array[T]{}
	self = append(self, items...)
	return self
}

func (self Array[T]) Size() int {
	return len(self)
}

func (self Array[T]) Empty() bool {
	return self.Size() == 0
}

func (self *Array[T]) Push(items ...T) {
	*self = append(*self, items...)
}

func (self *Array[T]) Pop() T {
	arr := *self
	v := arr[self.Size()-1]
	arr = arr[:self.Size()-1]
	*self = arr
	return v
}

func (self *Array[T]) UnShift(items ...T) {
	*self = append(items, *self...)
}

func (self *Array[T]) Shift() T {
	arr := *self
	v := arr[0]
	arr = arr[1:]
	*self = arr
	return v
}

func (self Array[T]) Sort(compare func(T, T) bool) {
	sort.Slice(self, func(i int, j int) bool {
		return compare(self[i], self[j])
	})
}

func (self Array[T]) Find(compare func(T) bool) int {
	return sort.Search(len(self), func(i int) bool {
		return compare(self[i])
	})
}
