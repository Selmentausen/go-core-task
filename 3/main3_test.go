package main

import (
	"reflect"
	"testing"
)

func TestStringIntMap_AddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("cherry", 100)
	m.Add("orange", 42)

	val, ok := m.Get("cherry")
	if !ok || val != 100 {
		t.Errorf("Expected 100 and true for 'cherry', got %d and %v", val, ok)
	}

	val, ok = m.Get("apple")
	if ok || val != 0 {
		t.Errorf("Expected 0 and false for 'apple', got %d and %v", val, ok)
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("test", 1)
	m.Remove("test")
	if m.Exists("test") {
		t.Errorf("Expected key 'test' to be removed, but Exists() returned true")
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("cherry", 100)
	if !m.Exists("cherry") {
		t.Errorf("Expected Exists to return true for key 'cherry'")
	}

	if m.Exists("apple") {
		t.Errorf("Expected Exists to return false for missing key 'apple'")
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	originalMap := NewStringIntMap()
	originalMap.Add("cherry", 100)
	originalMap.Add("orange", 42)

	copyMap := originalMap.Copy()

	expected := map[string]int{"cherry": 100, "orange": 42}
	if !reflect.DeepEqual(copyMap, expected) {
		t.Errorf("Copy failed. Expected %v, got %v", expected, copyMap)
	}

	// Verify independence
	copyMap["apple"] = 999
	val, _ := originalMap.Get("apple")
	if val == 999 {
		t.Errorf("Critical Error: Modifying the copy changed the original map!")
	}
}
