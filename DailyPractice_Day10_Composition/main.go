package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

type Parent struct {
	*Person
	ChildrenNumber int
}

type Parent2 struct {
	Person
	ChildrenNumber int
}

func main() {
	person := &Person{
		Name:    "Soualem Zahreddine",
		Surname: "Soualem Zahreddine",
		Age:     10,
	}
	person2 := Person{
		Name:    "Soualem Zahreddine",
		Surname: "Soualem Zahreddine",
		Age:     10,
	}

	parent := Parent{Person: person, ChildrenNumber: 4}
	parent2 := Parent2{Person: person2, ChildrenNumber: 4}

	//the struct
	/* fmt.Println(person)
	fmt.Println(person2) */

	//the composition
	fmt.Println(parent)
	fmt.Println(parent2)

	// you can either access the element by this two approaches
	/* fmt.Println(parent.Age)
	fmt.Println(parent.Person.Age) */

	// you can either access the element by this two approaches
	/* fmt.Println(parent2.Age)
	fmt.Println(parent2.Person.Age) */

	/* fmt.Println(parent.Person)
	fmt.Println(parent2.Person) */

	parent3 := CreateParent(person)
	fmt.Println(parent3)
	parent3.Age = 123
	fmt.Println(parent3.Age)
	fmt.Println(*parent3.Person)
	fmt.Println(parent3.Person)

	parent4 := CreateParent2(person2)
	parent4.Age = 123
	fmt.Println(parent4.Age)
	fmt.Println(parent4)

}

func CreateParent(person *Person) *Parent {
	return &Parent{
		Person:         person,
		ChildrenNumber: 5,
	}
}

func CreateParent2(person Person) Parent2 {
	return Parent2{
		Person:         person,
		ChildrenNumber: 5,
	}
}
