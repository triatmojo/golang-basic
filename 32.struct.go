package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct Method
func (c Customer) sayHeii(name string) {
	fmt.Println("Hello, My Name is ", c.Name)
}

func (c Customer) sayHuuu() {
	fmt.Println("Hello", c.Name)
}

func main() {

	// Step 1
	var brian Customer

	brian.Name = "Brian"
	brian.Address = "Bogor"
	brian.Age = 20

	fmt.Println(brian)

	// Struct Literal

	// Step 2 (Best Practice)
	john := Customer{
		Name:    "John",
		Address: "Bandung",
		Age:     20,
	}

	fmt.Println(john)

	// Step 3 (No Recommend)
	steve := Customer{"steve", "Solo", 20}
	fmt.Println(steve)

	// Struct method
	steve.sayHeii("jaka")
	john.sayHuuu()

}
