package main

import (
	"fmt"
	"time"
)

func main() {
	// 1
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 2
	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	// 3
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-11-16")
	fmt.Println(parse)

}
