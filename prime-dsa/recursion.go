package main

func add(n int) int {
	if n == 1 {
		return 1
	}
	out := n + add(n-1)
	return out
}
