package main

import "fmt"

func main() {

	var months = [...]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Append function in slice

	var slice2 = months[9:11]
	fmt.Println("slice 2:", slice2)

	var slice3 = append(slice2, "Hiyaaa")
	fmt.Println("slice 3: ", slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Make function slice

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Bruce"
	newSlice[1] = "Wayne"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy function slice

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// different make array and slice
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(array)
	fmt.Println(slice)
}
