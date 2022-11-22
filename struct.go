package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Bruce"
	customer.Address = "Gotham"
	customer.Age = 34
	customer.sayHai("Clark")

	customer2 := Customer{
		Name:    "Clark",
		Address: "Metropolis",
		Age:     30,
	}

	fmt.Println(customer.Name)
	fmt.Println(customer2.Name)
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
