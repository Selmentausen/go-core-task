package main

import (
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	limit := 100
	ch := RandomGenerator(limit)

	count := 0
	for val := range ch {
		count++
		if val < 0 {
			t.Errorf("Expected positive number, got %d", val)
		}
	}
	if count != limit {
		t.Errorf("Expected %d numbers, got %d", limit, count)
	}
}

func TestRandomGenerator_Zero(t *testing.T) {
	ch := RandomGenerator(0)
	count := 0
	for range ch {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0, got %d", count)
	}
}
