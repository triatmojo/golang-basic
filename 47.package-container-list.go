package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Steve")
	data.PushBack("Jobs")
	data.PushBack("Joyy")

	// fmt.Println(data.Front().Next().Next())

	// // fmt.Println(data.Back().Value)
	// // fmt.Println(data.Front().Value)

	// dari atas ke bawah
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari bawah ke atas
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
