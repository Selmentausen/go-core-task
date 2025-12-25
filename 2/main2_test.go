package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SliceExample returned %v, expected %v", result, expected)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := addElements(input, 11)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AddElements returned %v, expected %v", result, expected)
	}
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copyOfSlice := copySlice(originalSlice)
	if !reflect.DeepEqual(originalSlice, copyOfSlice) {
		t.Errorf("Copy should contain smae elemetns. returned %v, expected %v", originalSlice, copyOfSlice)
	}

	copyOfSlice[0] = 999
	if originalSlice[0] == 999 {
		t.Errorf("Modifying the copy changed the original slice!")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idxToRemove := 5
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	result := removeElement(input, idxToRemove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveElement returned %v, expected %v", result, expected)
	}

	resultSafe := removeElement(input, 100)
	if !reflect.DeepEqual(resultSafe, input) {
		t.Errorf("Expected original slice for out-of-bounds index, got %v", resultSafe)
	}
}
