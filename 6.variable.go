package main

import "fmt"

func main() {

	// step 1
	var name string

	name = "Tri"
	fmt.Println(name)

	// step 2
	var FriendName = "Tri Atmojo"
	fmt.Println(FriendName)

	var age = 20
	fmt.Println(age)

	// single variable

	country := "Indonesia"
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Tri"
		lastName  = "Atmojo"
	)

	fmt.Println(firstName, lastName)

}
