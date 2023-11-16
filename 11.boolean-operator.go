package main

import "fmt"

func main() {

	var (
		ujian      = 80
		absen      = 80
		lulusUjian = ujian >= 78
		lulusAbsen = absen >= 80
	)

	fmt.Println(lulusUjian && lulusAbsen)

	fmt.Println(ujian >= 78 && absen >= 80)
}
