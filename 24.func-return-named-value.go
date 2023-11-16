package main

import "fmt"

func getFullname() (firstName, lastName, address string) {
	firstName = "Brian"
	lastName = "Conner"
	address = "jakarta"

	return
}

func main() {
	a, b, c := getFullname()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
