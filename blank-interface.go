package main

import "fmt"

func Ups(i int) interface{} {
	if i%2 == 0 {
		return true
	} else if i%2 != 0 {
		return false
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)
}
