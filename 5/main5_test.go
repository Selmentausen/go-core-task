package main

import (
	"reflect"
	"testing"
)

func TestMapIntersection(t *testing.T) {
	tests := []struct {
		name          string
		a             []int
		b             []int
		expectedFound bool
		expectedSlice []int
	}{
		{
			name:          "Standard intersection",
			a:             []int{65, 3, 58, 678, 64},
			b:             []int{64, 2, 3, 43},
			expectedFound: true,
			expectedSlice: []int{64, 3},
		},
		{
			name:          "No intersection",
			a:             []int{1, 2, 3},
			b:             []int{4, 5, 6},
			expectedFound: false,
			expectedSlice: []int{},
		},
		{
			name:          "Full intersection",
			a:             []int{10, 20},
			b:             []int{10, 20},
			expectedFound: true,
			expectedSlice: []int{10, 20},
		},
		{
			name:          "One empty slice",
			a:             []int{},
			b:             []int{1, 2, 3},
			expectedFound: false,
			expectedSlice: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFound, gotSlice := MapIntersection(tt.a, tt.b)

			if gotFound != tt.expectedFound {
				t.Errorf("MapIntersection() found = %v, expected %v", gotFound, tt.expectedFound)
			}

			if len(gotSlice) == 0 && len(tt.expectedSlice) == 0 {
				return
			}

			if !reflect.DeepEqual(gotSlice, tt.expectedSlice) {
				t.Errorf("MapIntersection() slice = %v, expected %v", gotSlice, tt.expectedSlice)
			}
		})
	}
}
