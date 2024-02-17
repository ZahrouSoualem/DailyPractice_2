package main

import "fmt"

/*

This Go program defines a recursive function called MaxMin that takes a pointer to a slice of integers
as input and returns the maximum and minimum values in the slice. The main function creates a
test slice called table and calls MaxMin with a pointer to that slice. It then prints the result
of calling MaxMin, as well as the individual maximum and minimum values.

The MaxMin function uses a divide-and-conquer approach to recursively split the slice into two
halves, and then computes the maximum and minimum values for each half. It then returns
the maximum of the two maximum values and the minimum of the two minimum values.

*/

func main() {
	table := []int{18, 22, 37, 94, 50, 6, 11, 8, 15, 23, 7, 31, 43}
	n := int(len(table) / 2)

	fmt.Println(n)
	fmt.Println(table[:n])
	fmt.Println(table[n:])

	/* max, min := MaxMin(&table)
	fmt.Println("Max = ", max)
	fmt.Println("Min = ", min) */
}

func MaxMin(table *[]int) (int, int) {
	if len(*table) == 1 {
		return (*table)[0], (*table)[0]
	}
	n := int(len(*table) / 2)

	s1 := (*table)[:n]

	s2 := (*table)[n:]

	max1, min1 := MaxMin(&s1)
	max2, min2 := MaxMin(&s2)
	return Max(max1, max2), Min(min1, min2)
}

func Min(value1 int, value2 int) int {
	if value2 > value1 {
		return value1
	} else {
		return value2
	}
}

func Max(value1 int, value2 int) int {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}
