package main

import (
	"testing"
)

func TestPreOrderSearch(t *testing.T) {
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

	traversal := PreOrderSearch(*head)

	expected := []int{1, 2, 3, 4}

	if len(expected) != len(traversal) {
		t.Fatal("pre order traversal mismatched lengths")
	}
	for i, v := range expected {
		if v != expected[i] {
			t.Fatal("pre order traversal value mismatch")
		}
	}
}

func TestInOrderSearch(t *testing.T) {
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

	traversal := InOrderSearch(*head)

	expected := []int{3, 2, 1, 4}

	if len(expected) != len(traversal) {
		t.Fatal("in order traversal mismatched lengths")
	}
	for i, v := range expected {
		if v != expected[i] {
			t.Fatal("in order traversal value mismatch")
		}
	}
}

func TestPostOrderSearch(t *testing.T) {
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

	traversal := PostOrderSearch(*head)

	expected := []int{3, 2, 4, 1}

	if len(expected) != len(traversal) {
		t.Fatal("post order traversal mismatched lengths")
	}
	for i, v := range expected {
		if v != expected[i] {
			t.Fatal("post order traversal value mismatch")
		}
	}
}

func TestCompare(t *testing.T) {
	a := &BinaryNode{
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
	b := &BinaryNode{
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
	if Compare(a, b) != true {
		t.Fatal("tree compare")
	}
}
