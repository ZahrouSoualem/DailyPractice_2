package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	semaphore = make(chan struct{}, 5) // the semaphore with capacity 5
	wg        sync.WaitGroup           // the wait group for goroutines
)

func main() {
	// this is unbuffered channels
	/*
		The deadlock occurs in the given code because the channel channels is an unbuffered channel,
		meaning that it can only hold one value at a time. In this code, two values are being sent
		to the channel using the channels <- syntax. However, since the channel is unbuffered,
		the first send operation blocks until another goroutine reads from the channel.
		But in this case, there is no other goroutine that is attempting to read from the channel,
		so the program gets stuck in a deadlock state.
	*/
	/* channels := make(chan string)
	channels <- "Zahr Eddine Soualem"
	channels <- "Zahr Eddine Soualem"

	fmt.Println(<-channels) */

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i) // create a new worker goroutine
	}
	wg.Wait() // wait for all workers to finish
}

func worker(id int) {
	defer wg.Done() // signal to wait group that this worker is done
	fmt.Printf("Worker %d is waiting for the semaphore.\n", id)
	semaphore <- struct{}{} // acquire the semaphore
	fmt.Printf("Worker %d has acquired the semaphore.\n", id)
	time.Sleep(5 * time.Second) // do some work
	fmt.Printf("Worker %d is releasing the semaphore.\n", id)
	<-semaphore // release the semaphore
}
