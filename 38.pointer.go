package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Func Pointer
func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// Pointer => kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada

func main() {

	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Jakarta"

	// Pass by value
	// address2 := address1

	// Pass by reference
	// address2 := &address1

	// Operator
	*address2 = Address{"Malang", "Jawa tengah", "indonesia"}

	// New Pointer
	address4 := new(Address)
	address4.Country = "indonesia"

	address := Address{"Solo", "Jawa Timur", ""}
	changeAddressToIndonesia(&address)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address)
}
