package main

import (
	"fmt"
	"sort"
)

func main() {

	/*Simple Pointers*/

	a := 5
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)

	/*Table Pointers*/

	//Case One: Send a pointer
	table := []int{18, 22, 37, 94, 50, 6, 11, 8, 15, 23, 7, 31, 43}

	Aa(&table)
	fmt.Println("Sorted slice in ascending order: ", table)

	//Case Two: Send a value

	table = []int{18, 22, 37, 94, 50, 6, 11, 8, 15, 23, 7, 31, 43}

	Bb(table)
	fmt.Println("Sorted slice in ascending order: ", table)

	Clarification(&table)

}

func A(a int) {
	b := a
	a = 8 + b

}

func B(a *int) {
	*a = 8
}

func Aa(table *[]int) {
	sort.Ints(*table)
}

func Bb(table []int) {
	table = append(table, 12)
	sort.Ints(table)
}

func Clarification(table *[]int) {
	// Prints the slice of integers
	fmt.Println(table)
	/* The first fmt.Println(table) statement in the Clarification function will print the slice of integers as is,
	without dereferencing the pointer. */

	// Prints the memory address of the pointer to the slice of integers
	fmt.Println(&table)

	// Prints the contents of the slice of integers
	fmt.Println(*table)
	/* The third fmt.Println(*table) statement, on the other hand, will print the contents of the slice
	   of integers by dereferencing the pointer using the * operator. This means that it will print the actual
	    values contained in the slice, rather than the memory address of the pointer or the pointer itself. */
}
