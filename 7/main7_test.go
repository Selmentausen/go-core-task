package main

import (
	"sort"
	"testing"
)

func createChan(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func TestMergeChannels(t *testing.T) {
	ch1 := createChan(1, 2)
	ch2 := createChan(3, 4)
	ch3 := createChan(5)

	merged := MergeChannels(ch1, ch2, ch3)

	var results []int
	for n := range merged {
		results = append(results, n)
	}

	sort.Ints(results)
	expected := []int{1, 2, 3, 4, 5}

	if len(results) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(results))
	}

	for i, v := range expected {
		if results[i] != v {
			t.Errorf("Index %d: expected %d, got %d", i, v, results[i])
		}
	}
}

func TestMergeChannels_Empty(t *testing.T) {
	merged := MergeChannels()

	count := 0
	for range merged {
		count++
	}

	if count != 0 {
		t.Errorf("Expected 0 elements from empty merged, got %d", count)
	}
}
