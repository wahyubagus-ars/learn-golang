package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Bruce", spamFIlter)
}

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFIlter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
