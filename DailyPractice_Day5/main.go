package main

import (
	"fmt"
	"sync"
	_ "sync"
	"time"
)

var (
	a  int = 0
	mu sync.Mutex
)

func main() {

	go greeter(1)
	greeter(2)

}

func greeter(s int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		mu.Lock()
		fmt.Println("I am ", s, " before  a = ", a)
		a++
		fmt.Println("I am ", s, "after  a = ", a)
		mu.Unlock()
	}
}
