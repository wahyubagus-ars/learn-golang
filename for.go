package main

import "fmt"

func main() {
	counter := 1
	isContinue := false

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		if counter == 7 && !isContinue {
			break
		}
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke:", counter2)
	}

	slice := []string{"Senin", "Selasa", "Rabu"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For Range
	for index, name := range slice {
		fmt.Println("index:", index, "=", name)
	}
}
