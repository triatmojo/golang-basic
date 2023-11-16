package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(20, 30, 40, 50)
	fmt.Println(total)

	// Slice parameter => ketika ingin menambahkan parameter di func sumAll
	slice := []int{10, 30, 70, 60}
	total = sumAll(slice...)
	fmt.Println(total)
}
