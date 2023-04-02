package main

import (
	"fmt"
)

// data types
// user input
// Loops
// Conditional Statement
const variable = 2

func main() {
	// this is a declaration and assigning
	variable2 := variable
	fmt.Println(variable2)
	// this is a declaration with a variable and initialise it with a value
	var variable3 = variable
	fmt.Println(variable3)
	/* If the expression is omitt ed, the initial value is the zero value for the type,
	which is 0 for numbers, false for boole ans, "" for str ings, and nil
	for int erfaces and reference typ es (slice, point er, map, channel, function).  */
	var variable4 int
	fmt.Println(variable4)
	// User Input
	var number int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&number)
	fmt.Printf("the number is %d", number)
	m := map[string]int{
		"kaled": 9,
		"zaki":  10,
	}
	fmt.Println(m)
	delete(m, "kaled")
	fmt.Println(m)
	/* for key := range m {
		if key.expired() {
			delete(m, key)
		}
	} */
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	fmt.Println(sum)
	x := new(string)
	// even though we used new to allocate a memory but we did'nt initialize the value of type string ""
	fmt.Println(x)
	s := make([]int, 5)
	//  we used make to allocate a memoryand also  initialize the value of type string "" or int with 0 and so on
	fmt.Println(s)
}
