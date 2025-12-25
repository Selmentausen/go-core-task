package main

import (
	"fmt"
	"math"
)

func main() {
	inputCh := make(chan uint8)
	outputCh := make(chan float64)
	go RunPipeline(inputCh, outputCh)

	go func() {
		defer close(inputCh)
		numbers := []uint8{1, 2, 3, 5, 10}
		for _, n := range numbers {
			inputCh <- n
		}
	}()

	for result := range outputCh {
		fmt.Println(result)
	}
}

func RunPipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for num := range in {
		out <- math.Pow(float64(num), 3)
	}
}
