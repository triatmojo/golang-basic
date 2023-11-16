package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else if name == "B2" {
		return "Pig"
	} else {
		return name
	}

}

func main() {
	sayHelloWithFilter("Brian", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("B2", spamFilter)
}
