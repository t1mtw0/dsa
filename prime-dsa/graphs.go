package main

import (
	"math"
	"reflect"
)

type GraphEdge struct {
	to     int
	weight int
}
type WeightedAdjMatrix [][]int
type WeightedAdjList [][]GraphEdge

func BfsGraph(graph WeightedAdjMatrix, source, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	seen[source] = true

	var q DoublyLinkedList[int]
	q.Append(source)
	for ok := true; ok; ok = q.length != 0 {
		curr := q.Pop()
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Append(i)
		}
	}

	curr := needle
	var path []int
	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	if len(path) != 0 {
		path = append(path, source)
		reverse(path)
		return path
	}
	return []int{}
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func DfsGraph(graph WeightedAdjList, source, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	walk(graph, source, needle, seen, &path)

	if len(path) == 0 {
		return []int{}
	}

	return path
}

func walk(graph WeightedAdjList, curr, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true

	*path = append(*path, curr)

	if curr == needle {
		return true
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if walk(graph, edge.to, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func DijkstraList(graph WeightedAdjList, source, sink int) []int {
	// TODO: refactor to use MinHeap
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	dists := make([]float64, len(graph))
	for i := 0; i < len(dists); i++ {
		dists[i] = math.Inf(1)
	}
	dists[source] = 0.0

	for hasUnseen(dists, seen) {
		curr := getLowestUnseen(dists, seen)
		seen[curr] = true

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]

			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + float64(edge.weight)
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	curr := sink
	var path []int
	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	if len(path) != 0 {
		path = append(path, source)
		reverse(path)
		return path
	}
	return []int{}
}

func hasUnseen(dists []float64, seen []bool) bool {
	if len(dists) != len(seen) {
		return false
	}

	for i := 0; i < len(dists); i++ {
		if !seen[i] && dists[i] < math.Inf(1) {
			return true
		}
	}

	return false
}

func getLowestUnseen(dists []float64, seen []bool) int {
	idx := -1
	lowestDistance := math.Inf(1)

	if len(dists) != len(seen) {
		return idx
	}

	for i := 0; i < len(dists); i++ {
		if seen[i] {
			continue
		}
		if lowestDistance > dists[i] {
			idx = i
			lowestDistance = dists[i]
		}
	}

	return idx
}
