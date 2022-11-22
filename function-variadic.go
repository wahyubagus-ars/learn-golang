package main

import "fmt"

func main() {
	total := sumAll(2, 4, 6, 86, 43)

	// send slice as a variable args in variadic function
	numbers := []int{1, 2, 5, 4, 3}
	total2 := sumAll(numbers...)
	fmt.Println(total)
	fmt.Println(total2)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
