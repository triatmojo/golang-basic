package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	steve := Man{"steve"}
	steve.Married()

	fmt.Println(steve)
}
