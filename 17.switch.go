package main

import "fmt"

func main() {

	// Switch
	name := "brian"

	switch name {
	case "brian":
		fmt.Println("Halo, Brian")
	case "dani":
		fmt.Println("Halo, Dani")
	default:
		fmt.Println("Kenalan yok")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 10:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama suda benar")
	}
}
