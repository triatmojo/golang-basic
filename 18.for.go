package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for statement
	for counter := 1; counter <= 20; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for slice and array
	slice := []string{"Test1", "Test2", "Test3"}

	for index, value := range slice {
		fmt.Println("Index ke", index, "=", value)
	}

	// for range map
	person := make(map[string]string)
	person["name"] = "Brian"
	person["address"] = "Jakarta"

	for key, value := range person {
		fmt.Println(key, "", value)
	}

}
