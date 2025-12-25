package main

import "math/rand"

func RandomGenerator(limit int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < limit; i++ {
			ch <- rand.Intn(100)
		}
	}()
	return ch
}
