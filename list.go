package std

type ListNode[T any] struct {
	value *T
	prev  *ListNode[T]
	next  *ListNode[T]
}

type List[T any] struct {
	first *ListNode[T]
	last  *ListNode[T]
	size  int
}

func NewList[T any]() *List[T] {
	v := List[T]{}

	return &v
}

func (self *List[T]) Push(v T) *T {
	n := ListNode[T]{
		value: &v,
		prev: self.last,
	}

	self.last = &n

	if self.first == nil {
		self.first = &n
		self.first.next = nil
		self.first.prev = nil
	}

	self.size++
	return n.value
}

func (self *List[T]) Pop() *T {
	if self.last == nil {
		return nil
	}

	v := self.last.value
	self.last = self.last.prev

	return v
}

// func (self *List[T]) Remove(i int) *ListNode[T] {
// 	n := self.first

// 	for n != nil {


// 		n = n.next
// 	}
// }

func (self *List[T]) Size() int {
	return self.size
}
