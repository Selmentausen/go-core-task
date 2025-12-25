package main

import (
	"reflect"
	"testing"
)

func TestMapDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "Standard case with overlap",
			slice1:   []string{"apple", "banana", "cherry", "date"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry"},
		},
		{
			name:     "No overlap (keep all)",
			slice1:   []string{"one", "two"},
			slice2:   []string{"three", "four"},
			expected: []string{"one", "two"},
		},
		{
			name:     "Complete overlap (remove all)",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
		{
			name:     "Empty first slice",
			slice1:   []string{},
			slice2:   []string{"apple"},
			expected: []string{},
		},
		{
			name:     "Empty second slice",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{},
			expected: []string{"apple", "banana"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapDifference(test.slice1, test.slice2)
			if len(result) == 0 && len(test.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
