package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // biasakan menggunakan defer diatass
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result :", result)
	// contoh sebelum menggunakan defer dan memasukkan param value 0
	// maka akan terjadi error (panic)
	// logging() // maka ini tidak akan dieksekusi karena terjadi error
}

func main() {
	/**
	Defer :
	- adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	- defer function akan selalu dieksekusi walaupun terjadi error di function yg dieksekusi
	*/

	runApplication(10)
}
