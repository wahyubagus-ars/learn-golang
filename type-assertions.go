package main

import "fmt"

func random() interface{} {
	return 4
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	// switch case type assertions
	switch value := result.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	}
}
