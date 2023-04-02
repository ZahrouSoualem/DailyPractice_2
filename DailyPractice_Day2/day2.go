package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

// iota
// arrays
// slices
// functions and variadic functions
func main() {
	//Array
	list := [4]int{1, 2, 3, 4}
	var list2 [3]int = [3]int{1, 2, 3}
	var list3 = [3]int{1, 2, 3}
	var list4 [3]int

	fmt.Println(list)
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)

	/* In this case, the ellipsis notation is used to indicate that the
	size of the array should be determined by the number of elements
	in the initialization list. */

	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: ""}
	fmt.Println(RMB, symbol[RMB]) // "3 ""

	list5 := []int{1, 2, 3, 4, 5, 6, 7}
	var i int
	fmt.Scan(&i)
	for k, value := range list5 {
		if i == value {
			fmt.Println("the value is ", i)
			list5 = append(list5[:k], list5[k+1:]...)
		}
	}

	fmt.Println(list5)
	fmt.Println("here is the list ", list5[:])
	c := 0
	for i := 0; i < len(list5)/2; i++ {
		c = list5[len(list5)-i-1]
		list5[len(list5)-i-1] = list5[i]
		list5[i] = c
	}

	fmt.Println(list5)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	fmt.Println(s) // "[2 3 4 5 0 1]"
	reverse(s[2:])
	fmt.Println(s) // "[2 3 4 5 0 1]"
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	a1 := [5]int{1, 2, 3, 4, 5}
	s = a1[1:4]

	fmt.Println(len(s)) // prints 3
	fmt.Println(cap(s)) // prints 4
	fmt.Println(s[0], s[2])
	show(a1[:]...)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// variadic function
func show(nums ...int) {
	for _, value := range nums {
		fmt.Print(value, " ")
	}
}
