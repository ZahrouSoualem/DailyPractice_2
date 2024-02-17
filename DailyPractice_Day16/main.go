package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, cancel context.CancelFunc) {

	defer fmt.Println("the worer has ended his job succesfully")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("the worker is done")
			return
		default:
			fmt.Println("the worker is working")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, cancel)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
