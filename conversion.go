package main

import "fmt"

func main() {
	var nilai32 int = 127
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	fmt.Println("Conversion byte to string")

	name := "Bruce"
	e := name[0]
	eString := string(e)

	fmt.Println(e)
	fmt.Println(eString)
}
