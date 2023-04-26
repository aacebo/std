package list

type listNode[T any] struct {
	value T
	prev  *listNode[T]
	next  *listNode[T]
}

type List[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	size int
}

func New[T any]() List[T] {
	return List[T]{nil, nil, 0}
}

func (self List[T]) Size() int {
	return self.size
}

func (self List[T]) Empty() bool {
	return self.size == 0
}

func (self List[T]) First() T {
	return self.head.value
}

func (self List[T]) Last() T {
	return self.tail.value
}

func (self List[T]) At(i int) T {
	node := self.head
	j := 0

	for node != nil {
		if j == i {
			break
		}

		node = node.next
		j++
	}

	return node.value
}

func (self *List[T]) Push(items ...T) {
	for _, v := range items {
		node := listNode[T]{v, self.tail, nil}

		if self.tail != nil {
			self.tail.next = &node
		}

		self.tail = &node

		if self.head == nil {
			self.head = &node
			self.head.next = self.tail
		}

		self.size++
	}
}

func (self *List[T]) Pop() T {
	v := self.tail.value

	if self.tail.prev != nil {
		self.tail.prev.next = nil
	}

	self.tail = self.tail.prev
	self.size--
	return v
}

func (self *List[T]) UnShift(items ...T) {
	for i := len(items) - 1; i > -1; i-- {
		node := listNode[T]{items[i], nil, self.head}

		if self.head != nil {
			self.head.prev = &node
		}

		self.head = &node

		if self.tail == nil {
			self.tail = &node
			self.tail.prev = self.head
		}

		self.size++
	}
}

func (self *List[T]) Shift() T {
	v := self.head.value

	if self.head.next != nil {
		self.head.next.prev = nil
	}

	self.head = self.head.next
	self.size--
	return v
}

func (self *List[T]) Splice(i int, deleteCount int, items ...T) {
	node := self.head
	j := 0

	for node != nil {
		if j == i {
			break
		}

		node = node.next
		j++
	}

	k := 0

	for node != nil && k < deleteCount {
		if node.prev != nil {
			node.prev.next = node.next
		}

		if node.next != nil {
			node.next.prev = node.prev
		}

		if self.head == node {
			self.head = node.next
		}

		if self.tail == node {
			self.tail = node.prev
		}

		node = node.next
		self.size--
		k++
	}

	for _, v := range items {
		n := listNode[T]{v, nil, node}

		if node.prev != nil {
			n.prev = node.prev
			node.prev.next = &n
			node.prev = &n
		}

		self.size++
	}
}

func (self List[T]) ForEach(callback func(int, T)) {
	node := self.head
	i := 0

	for node != nil {
		callback(i, node.value)
		node = node.next
		i++
	}
}
