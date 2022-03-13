package main

import "fmt"

func main() {
	/*
		Tipe Data Map :
		- tipe data kumpulan key and value
		- key bersifat unik atau tidak boleh sama
		- jika memasukkan key yang sama, maka value sebelumnya akan diganti dengan data baru

		Function Map :
		- len(map)							=> Mendapatkan jumlah data di map
		- map[key]							=> Mengambil data di map menggunakan key
		- map[key] = value					=> Mengubah data di map menggunakan key
		- make(map[TypeData]TypeValue)		=> Membuat map baru
		- delete(map, key)					=> Mengahpus data di map dengan key
	*/

	person := map[string]string{
		"name":    "Fais Nasrullah",
		"address": "Depok, Jawa Barat",
	}

	// menambahkan key baru
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Membuat map baru menggunakan function make()
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Programmer Zaman Now"
	book["publisher"] = "Buku Zaman Now"
	book["wrong"] = "Oops"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
