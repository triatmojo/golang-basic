package main

import "fmt"

type Blaklist func(string) bool

func registerUser(name string, blaklist Blaklist) {
	if blaklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blaklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blaklist)
	registerUser("brian", blaklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

}
