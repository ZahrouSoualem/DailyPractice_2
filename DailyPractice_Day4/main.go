package main

import "fmt"

func main() {

	c1 := redisplugin.DefaultPlugin
	c2 := redisplugin.DefaultPlugin

	if c1 == c2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
