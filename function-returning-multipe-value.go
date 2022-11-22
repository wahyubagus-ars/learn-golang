package main

import "fmt"

func main() {
	firstname, lastname := getFullName()

	// only get spesific return value
	firstname2, _ := getFullName()
	fmt.Println(firstname, lastname)
	fmt.Println(firstname2)
}

func getFullName() (string, string) {
	return "Bruce", "Wayne"
}
