package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Steve Jobs", "Jobs"))
	fmt.Println(strings.Contains("Steve Jobs", "Nana"))

	fmt.Println(strings.Split("Steve Jobs", " "))

	fmt.Println(strings.ToUpper("Steve Jobs"))
	fmt.Println(strings.ToLower("Steve Jobs"))

	fmt.Println(strings.ReplaceAll("Jobs Jobs Jobs", "Jobs", "Budi"))

	fmt.Println(strings.Trim("   Stever Jobs   ", " "))

}
