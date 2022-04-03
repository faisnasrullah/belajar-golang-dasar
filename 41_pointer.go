package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	Pointer :
	- secara default di golang semua variabel itu di passing by value, bukan by reference
	- artinya jika kita mengirim sebuah variabel kedalam function, method atau variabel lain,
	  sebenernya yg dikirim adalah duplukasi valuenya
	- pointer adalah kemampuan membuat reference ke lokasi data di memory yg sama, tanpa menduplikasi data yg sudah ada
	- sederhananya dengan kemampuan pointer, kita bisa melakukan pass by reference
	*/

	// PASS BY VALUE
	address1 := Address{"Pekalongan", "Jawa Tengah", "Indonesia"}
	address2 := address1
	address2.City = "Depok" // meskipun disini address2 melakukan perubahan city menjadi depok, address1 tidak akan berubah

	fmt.Println(address1) // address1 tidak berubah (karena by default golang menggunakan PASS BY VALUE)
	fmt.Println(address2)
	fmt.Println("### MULAI IMPEMENTASI POINTER ###")

	// PASS BY REFERENCE (POINTER)
	/**
	- untuk membuat sebuah variabel dengan nilai pointer ke variabel yg lain
	  kita bisa menggunakan operator & lalu diikuti dengan nama variabelnya
	*/
	var address3 *Address = &address1 // address3 melakukan pointer ke variable address1
	address3.Country = "Singapore"    // address1 akan berubah sesuai perubahan yg dilakukan address3

	/**
	Operator *
	- siapapun yg mengacu ke lokasi memory yg sama, kalau kita menggunakan operator *
	  maka semuanya akan ikut berubah
	*/

	var address4 *Address = &address1
	*address4 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2) // address2 tidak ikut berubah, karena tidak melakukan pointer ke variabel address1
	fmt.Println(address3)
	fmt.Println(address4)

	/**
	Function New :
	- Cara lain membuat pointer dengan operator & adalah dengan menggunakan function new()
	- Namun function new() hanya akan mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/

	var address5 *Address = new(Address)
	address5.Province = "DKI Jakarta"
	fmt.Println(address5)
}
