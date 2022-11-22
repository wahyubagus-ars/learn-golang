package main

import "fmt"

func main() {
	sayHello()
	fmt.Println(sayHelloTo("Bruce", "Wayne"))
}

func sayHello() {
	fmt.Println("Say Hello")
}

func sayHelloTo(firstname string, lastname string) string {
	fmt.Println("Hello", firstname, lastname)

	return firstname + lastname
}
