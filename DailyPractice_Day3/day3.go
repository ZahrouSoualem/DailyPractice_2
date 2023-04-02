package main

import "fmt"

//maps
//struct
// interfaces

type A struct {
	name string
}

type B struct {
	a *A
}

func main() {

	//The make built-in function allocates and initializes an object of type slice, map, or chan (only).
	/* The new built-in function allocates memory. The first argument is a type, not a value,
	and the value returned is a pointer to a newly allocated zero value of that type. */
	myMpas := make(map[string]interface{})
	fmt.Println(myMpas)
	// this type is used to assign a value to a map after declaration to initialize the map
	myMpas["khaled"] = "soualem"
	myMpas["zahreddine"] = 25
	myMpas["Hello"] = "world"
	fmt.Println(myMpas)
	// this type is used only to initialize the map
	myMpas = map[string]interface{}{
		"holla": 2,
		"hola":  2,
		"hola2": 2,
	}
	fmt.Println(myMpas)

	a1 := A{
		name: "Soualem",
	}

	b1 := B{
		a: &a1,
	}
	fmt.Println(b1)
	fmt.Println(b1.a.name)
}
