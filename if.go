package main

import "fmt"

func main() {
	name := "bruce"

	if name == "bruce" {
		fmt.Println("bruce is a batman")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Name is valid")
	}

}
