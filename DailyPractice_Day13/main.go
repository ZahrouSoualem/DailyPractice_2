package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 5
	ch := make(chan int, 5)

	// Create 10 goroutines that write to the channel
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
			fmt.Printf("Goroutine %d wrote to channel\n", i)
		}(i)
	}

	// Wait for a short time to allow all goroutines to start writing
	time.Sleep(time.Second * 10)

	// Create a goroutine that tries to read from the channel
	for i := 0; i < 3; i++ {
		go func() {
			val := <-ch
			fmt.Printf("Goroutine read %d from channel\n", val)
		}()
	}

	// Wait for all goroutines to finish

	time.Sleep(time.Millisecond * 1000)

	fmt.Println("Program completed successfully")
}
