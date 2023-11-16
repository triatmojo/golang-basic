package main

import (
	"fmt"
	"golang-basic/database"

	// blank indentifier
	_ "golang-basic/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
