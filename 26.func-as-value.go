package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	// as value
	sayGoodBye := getGoodBye
	result := sayGoodBye("Brian")

	fmt.Println(result)
	fmt.Println(getGoodBye("Coner"))
}
