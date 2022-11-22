package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Master", blacklist)
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}
