package main

import "fmt"

func main() {
	firstname, middlename, lastname := getCompleteName()
	fmt.Println(firstname, middlename, lastname)
}

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "Master"
	middlename = "Bruce"
	lastname = "Wayne"

	// explicit declaration return value
	//return firstname, middlename, lastname

	// implicit declaration
	return
}
