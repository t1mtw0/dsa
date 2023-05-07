package main

import "testing"

func TestNewLRUCache(t *testing.T) {
	lru := NewLRUCache[string, int](10)
}

func TestLRUUpdate(t *testing.T) {
	lru := NewLRUCache[string, int](10)
}

func TestLRUGet(t *testing.T) {
	lru := NewLRUCache[string, int](10)
}

func TestLRUEvict(t *testing.T) {
	lru := NewLRUCache[string, int](10)
}
