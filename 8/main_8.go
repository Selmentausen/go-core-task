package main

import (
	"sync"
)

type CustomWaitGroup struct {
	counter int
	waitCh  chan struct{}
	mu      sync.Mutex
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		counter: 0,
		waitCh:  nil,
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	if wg.counter == 0 && delta > 0 {
		wg.waitCh = make(chan struct{})
	}

	wg.counter += delta
	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.counter == 0 && wg.waitCh != nil {
		close(wg.waitCh)
		wg.waitCh = nil
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	currentCh := wg.waitCh
	currentCounter := wg.counter
	wg.mu.Unlock()
	if currentCounter == 0 || currentCh == nil {
		return
	}
	<-currentCh
}
