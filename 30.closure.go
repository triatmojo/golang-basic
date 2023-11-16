package main

import "fmt"

// closure => sebuah func atau blok scope yang dapat berinteraksi dengan scope lainnya

func main() {

	counter := 0
	name := "Brian"
	increment := func() {
		name = "dani"
		name := "Rama"
		fmt.Println("increment")
		fmt.Println(name)
		counter++
	}

	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
