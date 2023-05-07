package main

import "testing"

func TestBfs(t *testing.T) {
	head := &BinaryNode{
		value: 1,
		left: &BinaryNode{
			value: 2,
			left: &BinaryNode{
				value: 3,
			},
			right: nil,
		},
		right: &BinaryNode{
			value: 4,
			left:  nil,
			right: nil,
		},
	}

	if Bfs(*head, 1) != true {
		t.Fatal("breadth first search first item")
	}
	if Bfs(*head, 2) != true {
		t.Fatal("breadth first search second item")
	}
	if Bfs(*head, 5) != false {
		t.Fatal("breadth first search item does not exist")
	}
}
