package main

import "fmt"

func main() {

	// break
	for i := 0; i < 20; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// continue
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}

}
