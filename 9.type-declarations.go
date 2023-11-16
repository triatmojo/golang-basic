package main

import "fmt"

func main() {
	// Membuat alias
	type (
		country string
		marrige bool
	)

	var (
		countryTri    country = "Indonesia"
		marrigeTri    marrige = false
		countryIchsan country = "Korea"
		marrigeIchsan marrige = true
	)

	fmt.Println(countryTri, marrigeTri)
	fmt.Println(countryIchsan, marrigeIchsan)
}
