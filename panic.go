package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	message := recover()
	fmt.Println("Error dengan msg", message)
	fmt.Println("Application selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error")
	}

	fmt.Println("Aplikasi Berjalan")
}
