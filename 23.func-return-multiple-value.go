package main

import "fmt"

func getPerson() (string, string, string) {
	return "Brian", "Corner", "Jakarta"
}

func main() {
	firstName, _, address := getPerson()
	fmt.Println(firstName)
	fmt.Println(address)
}
