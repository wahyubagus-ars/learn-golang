package main

import "fmt"

func main() {
	goodbye := getGoodBye

	fmt.Println(goodbye("Alfred"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
