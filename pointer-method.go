package main

import "fmt"

type Men struct {
	Name string
}

func (men *Men) Married() {
	men.Name = "Mr. " + men.Name
}

func main() {
	men := Men{"Bruce"}
	men.Married()

	fmt.Println(men)
}
