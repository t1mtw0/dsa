package main

import (
	"testing"
)

func TestBfsGraph1(t *testing.T) {
	adjMatrix := WeightedAdjMatrix{
		{0, 1, 5, 4},
		{1, 0, 0, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
	}
	path := BfsGraph(adjMatrix, 0, 2)
	expected := []int{0, 2}

	if len(path) != len(expected) {
		t.Fatal("bfs graph weighted adjacency matrix length mismatch")
	}
	for i := range path {
		if path[i] != expected[i] {
			t.Fatal("bfs graph weighted adjacency matrix")
		}
	}
}

func TestBfsGraph2(t *testing.T) {
	adjMatrix := WeightedAdjMatrix{
		{0, 3, 2, 0},
		{1, 0, 0, 0},
		{5, 4, 0, 0},
		{1, 2, 3, 0},
	}
	path := BfsGraph(adjMatrix, 0, 3)
	expect := []int{}

	if len(path) != len(expect) {
		t.Fatal("bfs graph weighted adjacency matrix not empty")
	}
}

func TestDfsGraph(t *testing.T) {
	adjList := WeightedAdjList{
		{GraphEdge{to: 1, weight: 1}, GraphEdge{to: 2, weight: 5}, GraphEdge{to: 3, weight: 4}},
		{GraphEdge{to: 0, weight: 1}},
		{GraphEdge{to: 3, weight: 2}},
		{},
	}
	path := DfsGraph(adjList, 0, 2)
	expected := []int{0, 2}

	if len(path) != len(expected) {
		t.Fatal("dfs graph weighted adjacency list length mismatch")
	}
	for i := range path {
		if path[i] != expected[i] {
			t.Fatal("dfs graph weighted adjacency list")
		}
	}
}

func TestDijkstraList(t *testing.T) {
	adjList := WeightedAdjList{
		{GraphEdge{to: 1, weight: 5}, GraphEdge{to: 2, weight: 2}, GraphEdge{to: 3, weight: 6}},
		{GraphEdge{to: 4, weight: 2}},
		{GraphEdge{to: 3, weight: 8}},
		{GraphEdge{to: 1, weight: 4}},
		{},
	}
	path := DijkstraList(adjList, 0, 4)
	expected := []int{0, 1, 4}

	if len(path) != len(expected) {
		t.Fatal("dijkstra weighted adjacency list length mismatch")
	}
	for i := range path {
		if path[i] != expected[i] {
			t.Fatal("dijkstra weighted adjacency list")
		}
	}
}
