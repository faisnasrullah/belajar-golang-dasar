package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// tambahkan operator * untuk memberi tahu bahwa parameter ini adalah pointer
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	/**
	Pointer di Function :
	- saat kita membuat parameter di function, by default adalah pass by value
	  artinya data akan di copy lalu dikirimkan ke function tersebut
	- oleh karena itu, jika kita mengubah data didalam function, data aslinya tidak akan pernah berubah
	- hal ini membuat variabel menjadi aman, karena tidak akan bisa diubah
	- namun ada case tertentu kita ingin membuat function yang bisa mengubah data asli parameter tersebut
	- untuk melakukan hal tersebut, kita bisa menggunakan pointer di function
	- untuk menjadikan sebuah parameter menjadi pointer, kita bisa menggunakan operator * pada parameternya
	*/

	// sebelum implementasi pointer ke dalam function
	// tidak terjadi perubahan pada variabel alamat
	alamat := Address{"Pekalongan", "Jawa Tengah", ""}

	// kita coba implementasi pointer
	var alamatPointer *Address = &alamat
	ChangeAddressToIndonesia(alamatPointer)

	fmt.Println(alamat)
}
