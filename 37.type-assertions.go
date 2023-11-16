package main

import "fmt"

func random() interface{} {
	return false
}

// type assertions => merubah tipe data menjadi tipe data yang diinginkan

func main() {
	result := random()
	// resultToString := result.(string)
	// fmt.Println(resultToString)

	// Using Switch Expression
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}

}
