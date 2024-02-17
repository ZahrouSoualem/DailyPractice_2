package main

import "fmt"

func main() {
	channels := make(chan string, 2)
	channels <- "Zahr Eddine Soualem"
	channels <- "Zahr Eddine Soualem2"

	fmt.Println(<-channels)
	channels <- "Zahr Eddine Soualem"

	fmt.Println(<-channels)
}
