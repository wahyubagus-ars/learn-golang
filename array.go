package main

import "fmt"

func main() {

	// create array with manual way
	var names [3]string
	names[0] = "Wahyu"
	names[1] = "Bagus"
	names[2] = "Sulaksono"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// create array with assigment
	var values = [3]int{
		90,
		93,
		35,
	}

	fmt.Println(values)

	//get length array (note: not length data inside of array)
	fmt.Println(len(values))
}
