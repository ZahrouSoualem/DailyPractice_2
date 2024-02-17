package main

import "fmt"

func main() {

	c1 := redisplug.DefaultPlugin
	c2 := redisplug.DefaultPlugin

	if c1 == c2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
