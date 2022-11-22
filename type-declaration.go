package main

import "fmt"

func main() {
	type NoKTP string

	var ktpNumber NoKTP = "124029391004"
	fmt.Println(ktpNumber)
	fmt.Println(NoKTP("234112424"))
}
