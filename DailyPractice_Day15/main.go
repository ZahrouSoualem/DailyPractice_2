package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	size  int
	queue chan struct{}
	wg    sync.WaitGroup
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		size:  size,
		queue: make(chan struct{}, size),
	}
}

func (s *Semaphore) Acquire() {
	s.queue <- struct{}{}
	s.wg.Add(1)
}

func (s *Semaphore) Release() {
	<-s.queue
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func main() {
	
	sem := NewSemaphore(3)

	for i := 0; i < 10; i++ {
		sem.Acquire()
		go func(i int) {
			fmt.Printf("Acquired semaphore for %d\n", i)
			sem.Release()
		}(i)
	}

	sem.Wait()
	fmt.Println("All goroutines have completed")
}
