package main

import "fmt"

func main() {
	name := "nana"

	if name == "Tri" {
		fmt.Println("Hello ", name)
	} else if name == "nana" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Nama salah!!!")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Data terlalu panjang")
	} else {
		fmt.Println("Data sudah benar")
	}
}
