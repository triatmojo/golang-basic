package main

import "fmt"

type Hasname interface {
	GetName() string
}

func sayHello(hasname Hasname) {
	fmt.Println("Hello", hasname.GetName())
}

// 1
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// 2
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var brian Person
	brian.Name = "brian"
	sayHello(brian)

	cat := Animal{Name: "Kucing"}
	fmt.Println(cat)
}
