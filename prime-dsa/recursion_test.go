package main

import "testing"

func TestRecursion(t *testing.T) {
	sum := add(5)
	if sum != 15 {
		t.Fatal("recursion sum")
	}
}
