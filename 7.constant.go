package main

import "fmt"

func main() {
	const firstName string = "Tri"
	const lastName = "Atmojo"

	fmt.Println(firstName, lastName)

	// multiple constant
	const (
		address string = "Jl.Iskandar"
		phone          = "08976617271"
	)

	fmt.Println(address)
	fmt.Println(phone)
}
