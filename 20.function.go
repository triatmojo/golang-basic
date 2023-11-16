package main

import "fmt"

func person() {
	person := map[string]string{
		"name":    "Brian",
		"address": "jakarta",
		"sport":   "futsal",
	}
	fmt.Println(person)
	fmt.Println(len(person))
}

func main() {
	person()
}
