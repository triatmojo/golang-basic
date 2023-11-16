package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello brother"
	}

	return "Hello " + name
}

func main() {

	result := getHello("Brian")

	fmt.Println(result)
	fmt.Println(getHello(""))
}
