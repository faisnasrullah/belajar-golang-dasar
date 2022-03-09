package main

import "fmt"

func main() {
	/*
		Constant :
		- variabel yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
		- gunakan kata kunci const
		- wajib langsung menginisialisasikan nilainya
	*/

	const (
		firstName string = "Fais"
		lastName         = "Nasrullah"
	)

	const discount = 7500

	// akan error jika kita coba re assign value nya
	// firstName = "Eko"
	// lastName = "Sungkowo"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(discount)
}
