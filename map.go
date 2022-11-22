package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Bruce Wayne",
		"age":  "34",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	person["address"] = "Gotham"
	fmt.Println(person["address"])

	// Make function

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Batara"
	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)
}
