package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("originalSlice:", originalSlice)
	evenSlice := sliceExample(originalSlice)
	fmt.Println("evenSlice:", evenSlice)
	addSlice := addElements(originalSlice, 11)
	fmt.Println("addSlice:", addSlice)
	newSlice := copySlice(originalSlice)
	fmt.Println("newSlice:", newSlice)
	removeSlice := removeElement(originalSlice, 5)
	fmt.Println("removeSlice:", removeSlice)
	fmt.Println("originalSlice:", originalSlice)
}

func sliceExample(originalSlice []int) []int {
	evenSlice := make([]int, 0)
	for _, v := range originalSlice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

func addElements(originalSlice []int, value int) []int {
	newSlice := make([]int, len(originalSlice))
	copy(newSlice, originalSlice)
	return append(newSlice, value)
}

func copySlice(originalSlice []int) []int {
	newSlice := make([]int, len(originalSlice))
	copy(newSlice, originalSlice)
	return newSlice
}

func removeElement(originalSlice []int, idx int) []int {
	if idx < 0 || idx >= len(originalSlice) {
		return originalSlice
	}
	newSlice := make([]int, 0, len(originalSlice)-1)
	newSlice = append(newSlice, originalSlice[:idx]...)
	newSlice = append(newSlice, originalSlice[idx+1:]...)
	return newSlice
}
