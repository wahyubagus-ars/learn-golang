package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, error := Pembagian(6, 0)
	if error == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error:", error.Error())
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagaian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}
