package main

import "fmt"

func main() {
	list := []int{10, 3, 7, 1, 9, 4, 8, 2, 6, 5}
	fmt.Println("From this unsorted list : ", list)
	list2 := MergeSort(&list)
	fmt.Println("To this sorted list     : ", list2)

}

func MergeSort(list1 *[]int) []int {
	s1 := *list1
	if len(s1) == 1 {
		return s1
	}
	mid := len(s1) / 2
	s2 := s1[:mid]
	s3 := s1[mid:]
	s2 = MergeSort(&s2)
	s3 = MergeSort(&s3)
	s4 := Merge(&s2, &s3)
	return s4
}

func Merge(list1 *[]int, list2 *[]int) []int {
	i, j := 0, 0
	list3 := []int{}

	s1 := *list1
	s2 := *list2
	for i < len(*list1) && j < len(*list2) {
		if s1[i] < s2[j] {
			list3 = append(list3, s1[i])
			i++
		} else {
			list3 = append(list3, s2[j])
			j++
		}
	}
	for i < len(*list1) {
		list3 = append(list3, s1[i])
		i++
	}

	for j < len(*list2) {
		list3 = append(list3, s2[j])
		j++
	}

	return list3
}
