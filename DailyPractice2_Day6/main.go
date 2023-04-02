package main

import (
	"fmt"
	"sync"
	"time"
)

var a = 0

func main() {
	//var wg sync.WaitGroup

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("I am ", 1, " before  a = ", a)
		a++
		fmt.Println("I am ", 1, "after  a = ", a)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("I am ", 2, " before  a = ", a)
		a++
		fmt.Println("I am ", 2, "after  a = ", a)

	}()

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("I am ", 3, " before  a = ", a)
		a++
		fmt.Println("I am ", 3, "after  a = ", a)

	}()

	wg.Wait()

	fmt.Println("all goroutines completed")
}

/* func greeter(s int) {
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second)
		mu.Lock()
		fmt.Println("I am ", s, " before  a = ", a)
		a++
		fmt.Println("I am ", s, "after  a = ", a)
		mu.Unlock()
	}
} */
