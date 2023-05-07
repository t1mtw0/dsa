package main

import (
	"math"
)

func binarysearch(haystack []int, low int, high int, n int) bool {
	for low < high {
		m := int(math.Floor(float64(low) + float64(high-low)/2.0))
		v := haystack[m]
		if v == n {
			return true
		} else if v > n {
			high = m
		} else {
			low = m + 1
		}
	}
	return false
}

func twocrystalballs(breaks []bool) int {
	jump_amount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	var i int
	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jump_amount
	}

	for j := i - jump_amount; j < i; j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
