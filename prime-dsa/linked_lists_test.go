package main

import (
	"testing"
)

func TestLLPrepend(t *testing.T) {
	var ll DoublyLinkedList[int]
	for i := 2; i >= 1; i-- {
		ll.Prepend(i)
	}

	if ll.Get(0) != 1 {
		t.Fatal("first value")
	}
	if ll.Get(1) != 2 {
		t.Fatal("second value")
	}
}

func TestLLInsert(t *testing.T) {
	var ll DoublyLinkedList[int]

	ll.Insert(1, 0)
	ll.Insert(2, 1)

	if ll.Get(0) != 1 {
		t.Fatal("first value")
	}
	if ll.Get(1) != 2 {
		t.Fatal("second value")
	}
}

func TestLLRemoveAt(t *testing.T) {
	var ll DoublyLinkedList[int]

	ll.Prepend(0)
	ll.Prepend(1)
	ll.Prepend(2)

	ll.RemoveAt(1)
	if ll.Get(1) != 0 {
		t.Fatal("remove at one value")
	}
	ll.RemoveAt(1)
	if ll.Get(0) != 2 {
		t.Fatal("remove at two values")
	}
}

func TestLLRemove(t *testing.T) {
	var ll DoublyLinkedList[int]

	ll.Append(0)
	ll.Append(1)
	ll.Append(2)

	ll.Remove()
	if ll.Get(1) != 1 {
		t.Fatal("remove one value")
	}
	ll.Remove()
	if ll.Get(0) != 0 {
		t.Fatal("remove two values")
	}
}

func TestLLAppend(t *testing.T) {
	var ll DoublyLinkedList[int]

	ll.Append(0)
	ll.Append(1)

	if ll.Get(0) != 0 {
		t.Fatal("append first value")
	}
	if ll.Get(1) != 1 {
		t.Fatal("append second value")
	}
}

func TestLLPop(t *testing.T) {
	var ll DoublyLinkedList[int]

	ll.Prepend(0)
	ll.Prepend(1)

	first := ll.Pop()
	if first != 1 {
		t.Fatal("pop value")
	}
	second := ll.Pop()
	if second != 0 {
		t.Fatal("pop second value")
	}
}
