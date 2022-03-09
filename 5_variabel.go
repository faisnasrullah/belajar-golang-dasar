package main

import "fmt"

func main() {
	/*
		Variabel pada Go :
		- Hanya bisa menyimpan tipe data yang sama
		- Untuk membuat variabel, kita harus menggunakan kata kunci var, lalu diikuti nama variabel dan tipe datanya
	*/
	// cara pertama
	var name string

	// cara kedua (golang akan secara otomatis tau tipe datanya apa)
	var email = "fais@automatalabs.co"
	var age byte = 23

	// cara ketiga (tanpa menggunakan kata kunci var)
	country := "Indonesia"

	name = "Fais Nasrullah"
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
	fmt.Println(country)

	name = "Eko Sungkowo"
	country = "Singapore"
	fmt.Println(name)
	fmt.Println(country)
}
