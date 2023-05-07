package main

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

func (d *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}
	d.length++
	if d.head == nil {
		d.head = newNode
		return
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}
	d.length++
	if d.head == nil {
		d.head = newNode
		return
	}
	n := d.head
	for n.next != nil {
		n = n.next
	}
	n.next = newNode
	newNode.prev = n
}

func (d *DoublyLinkedList[T]) Insert(value T, index int) {
	newNode := &Node[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}
	n := d.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	d.length++
	if d.head == nil {
		d.head = newNode
		return
	}
	if d.length-1 == index {
		d.Append(value)
		return
	}
	if d.length-1 > index {
		return
	}
	if index == 0 {
		d.Prepend(value)
		return
	}
	newNode.next = n
	n.prev.next = newNode
	newNode.prev = n.prev
	n.prev = newNode
}

func (d *DoublyLinkedList[T]) Remove() {
	if d.head == nil {
		return
	}
	n := d.head
	for i := 0; i < d.length-1; i++ {
		n = n.next
	}
	d.length--
	n.next = nil
}

func (d *DoublyLinkedList[T]) RemoveAt(index int) {
	if d.head == nil {
		return
	}
	if index == 0 {
		d.Pop()
		return
	}
	if index == d.length-1 {
		d.Remove()
		return
	}
	if index > d.length-1 {
		return
	}
	n := d.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	d.length--
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (d *DoublyLinkedList[T]) Pop() T {
	if d.head == nil {
		var empty T
		return empty
	}
	d.length--
	head := d.head
	d.head = head.next
	return head.value
}

func (d *DoublyLinkedList[T]) Get(index int) T {
	if d.head == nil {
		var empty T
		return empty
	}
	if index > d.length-1 {
		var empty T
		return empty
	}
	n := d.head
	for i := 0; i < index; i++ {
		if n.next != nil {
			n = n.next
		}
	}
	return n.value
}
