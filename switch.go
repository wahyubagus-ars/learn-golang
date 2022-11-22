package main

import "fmt"

func main() {
	name := "bruce"

	switch name {
	case "bruce":
		fmt.Println("Hallo bruce")
	case "wayne":
		fmt.Println("Hello master")
	default:
		fmt.Println("Hey")
	}

	// switch short statement

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Name too long")
	//case false:
	//	fmt.Println("Great Name")
	//}

	// switch without condition
	length2 := len(name)
	switch {
	case length2 > 5:
		fmt.Println("Name too long")
	case length2 <= 5:
		fmt.Println("Great Name")
	}
}
