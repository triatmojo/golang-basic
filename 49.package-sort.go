package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	user := []User{
		{"nana", 10},
		{"nini", 30},
		{"nunu", 50},
		{"nunu", 20},
		{"nunu", 40},
	}

	fmt.Println(user)
	sort.Sort(UserSlice(user))
	fmt.Println(user)
}
