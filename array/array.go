package array

import (
	"sort"
)

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

func (self Array[T]) First() T {
	return self[0]
}

func (self Array[T]) Last() T {
	return self[self.Size()-1]
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

func (self *Array[T]) Splice(i int, deleteCount int, items ...T) {
	arr := *self
	before := arr[:i]
	after := arr[i:]

	if deleteCount > 0 {
		after = arr[i+1:]
	}

	arr = append(before, append(items, after...)...)
	*self = arr
}

func (self Array[T]) Sort(compare func(T, T) bool) {
	sort.Slice(self, func(i int, j int) bool {
		return compare(self[i], self[j])
	})
}

func (self Array[T]) Find(compare func(T) int) int {
	left, right := 0, self.Size()-1

	for left <= right {
		mid := (left + right) / 2
		dif := compare(self[mid])

		if dif == 0 {
			return mid
		} else if dif < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (self Array[T]) Filter(compare func(T) bool) Array[T] {
	arr := Array[T]{}

	for _, v := range self {
		if compare(v) {
			arr = append(arr, v)
		}
	}

	return arr
}

func (self Array[T]) Slice(start int, end int) Array[T] {
	return self[start:end]
}
