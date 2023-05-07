package main

type lRU[K comparable, V any] struct {
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
	length        int
	capacity      int
}

func NewLRUCache[K comparable, V any](capacity int) lRU[K, V] {
	var lru lRU[K, V]
	lru.lookup = make(map[K]*Node[V], capacity)
	lru.reverseLookup = make(map[*Node[V]]K, capacity)
	lru.capacity = capacity
	return lru
}

func (l *lRU[K, V]) Update(key K, value V) {
	node, ok := l.lookup[key]
	if !ok {
		node := &Node[V]{
			value: value,
			next:  nil,
			prev:  nil,
		}
		l.length++
		l.prepend(node)
		l.trimCache()
		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.value = value
	}

}

func (l *lRU[K, V]) Get(key K) V {
	node, ok := l.lookup[key]
	if !ok {
		var empty V
		return empty
	}

	l.detach(node)
	l.prepend(node)

	return node.value
}

func (l *lRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = node.next
	}

	if l.tail == node {
		l.tail = node.prev
	}

	node.next = nil
	node.prev = nil

}

func (l *lRU[K, V]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *lRU[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}
