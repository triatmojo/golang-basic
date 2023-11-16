package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("Perkenalkan sama saya", firstName, lastName, "berusia", age, "tahun")
}

func main() {
	sayHelloTo("Brian", "Conner", 30)
}
