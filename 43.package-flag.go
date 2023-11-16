package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your database user")
	pass := flag.String("pass", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your database number")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("Username: ", *user)
	fmt.Println("Password: ", *pass)
	fmt.Println("Password: ", *number)
}
