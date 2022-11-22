package main

import "fmt"

func main() {
	const firstname = "Bruce"
	const lastName = "Wayne"

	fmt.Println(firstname)
	fmt.Println(lastName)

	// multipe constant
	const (
		name = "Dyana"
		age  = 5000
	)

	fmt.Println("name:", name, "age:", age)
}
