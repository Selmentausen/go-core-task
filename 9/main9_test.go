package main

import (
	"reflect"
	"testing"
)

func TestRunPipeline(t *testing.T) {
	inputCh := make(chan uint8)
	outputCh := make(chan float64)

	go RunPipeline(inputCh, outputCh)

	input := []uint8{1, 2, 3, 5, 10}
	expected := []float64{1, 8, 27, 125, 1000}

	go func() {
		defer close(inputCh)
		for _, val := range input {
			inputCh <- val
		}
	}()

	var result []float64
	for val := range outputCh {
		result = append(result, val)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRunPipelineEmpty(t *testing.T) {
	inputCh := make(chan uint8)
	outputCh := make(chan float64)
	go RunPipeline(inputCh, outputCh)

	close(inputCh)

	count := 0
	for range outputCh {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 results from empty input, got %d", count)
	}
}
