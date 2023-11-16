package main

import "fmt"

// Defer => function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
func looging() {
	fmt.Println("Selesai memanggil function")
}

func RunApplication(value int) {
	defer looging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Hasil", result)
}

// Panic => digunakan untuk menghentikan program
//
//	=> biasa digunakan ketika error pada saat kode program kita berjalan
//
// Recover => digunakan untuk menangkap data panic
func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	RunApplication(10)
	runApp(true)
}
