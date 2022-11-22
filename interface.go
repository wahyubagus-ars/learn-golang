package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type HashName interface {
	GetName() string
}

func main() {
	var person Person
	person.Name = "Bruce"

	sayHello2(person)

	animal := Animal{
		Name: "Bat",
	}

	sayHello2(animal)
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello2(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}
