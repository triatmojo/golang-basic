package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Tri",
		"address": "Bogor",
	}

	person["title"] = "DevOps"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Print(len(person))

	book := make(map[string]string)
	book["title"] = "Golang Pemula"
	book["author"] = "Tri"
	book["message"] = "ERROR"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "message")

	fmt.Println(book)
	fmt.Println(len(book))

	// Test
	ktp := make(map[int]int)
	ktp[1] = 9891827319321
	ktp[2] = 9891827319321

	fmt.Println(ktp)
}
