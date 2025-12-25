package main

import (
	"sync"
	"testing"
	"time"
)

func TestBasicWorkflow(t *testing.T) {
	wg := NewCustomWaitGroup()
	finished := false
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		finished = true
	}()

	wg.Wait()

	if !finished {
		t.Error("Wait() returned before worker finished")
	}
}

func TestReuse(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Done()
	}()
	wg.Wait()

	completeChan := make(chan bool)
	wg.Add(2)

	for range 2 {
		go func() {
			defer wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		completeChan <- true
	}()

	select {
	case <-completeChan:
		// Success
	case <-time.After(100 * time.Millisecond):
		t.Error("Round 2: Wait() timed out, likely reuse logic failed")
	}
}

func TestStress(t *testing.T) {
	wg := NewCustomWaitGroup()
	const workers = 100
	var counter int
	var mu sync.Mutex
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	if counter != workers {
		t.Errorf("Expected %d workers to finish, got %d", workers, counter)
	}
}

func TestNegativePanic(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic on negative counter")
		}
	}()

	wg.Add(1)
	wg.Done()
	wg.Done()
}

func TestZeroWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		// Success
	case <-time.After(50 * time.Millisecond):
		t.Error("Wait() blocked on empty WaitGroup")
	}
}
