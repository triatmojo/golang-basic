package main

import "fmt"

func test(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return false
	} else {
		return "Wrong"
	}
}

func main() {
	var data = test(1)
	fmt.Println(data)
}
