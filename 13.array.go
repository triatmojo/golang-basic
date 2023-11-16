package main

import "fmt"

func main() {

	var country [4]string
	country[0] = "Indonesia"
	country[1] = "Malaysia"
	country[2] = "Thailand"
	country[3] = "Japan"

	fmt.Println(country[0])
	fmt.Println(country[1])
	fmt.Println(country[2])
	fmt.Println(country[3])

	var value = [4]int{
		70,
		80,
		90,
		100,
	}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])
	fmt.Println(value[3])

	// len
	var test [10]int
	fmt.Println(len(country))
	fmt.Println(len(value))
	fmt.Println(len(test))
}
