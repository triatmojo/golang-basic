package helper

import "fmt"

var Address = "Jl.Setia"

func SayHello(name string) {
	name = "Hello " + name

	fmt.Println(name)
}
