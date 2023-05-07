package main

import "testing"

func TestBinarySearch(t *testing.T) {
	res := binarysearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 9, 0)
	if res != false {
		t.Fatal("not in array")
	}
}

func TestTwoCrystalBalls(t *testing.T) {
	res := twocrystalballs([]bool{false, false, false, false, false, false, false, false, true})
	if res != 8 {
		t.Fatal("list of length 9 last element")
	}
}
