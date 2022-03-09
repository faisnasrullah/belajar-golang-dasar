package main

import "fmt"

func main() {
	/*
		Tipe Data String :
		- direpresentasikan dengan kata kunci string
		- selalu diawali dan diakhiri dengan karakter " (petik dua)

		Function untuk String :
		- len("string")		=> Menghitung jumlah karakter pada tipe data string
		- "string"[number]	=> Mengambil karakter pada posisi yang ditentukan
	*/

	fmt.Println("Belajar Go Lang Dasar")
	fmt.Println(len("Belajar Go Lang Dasar"))
	fmt.Println("Belajar Go Lang Dasar"[4]) // return byte dari huruf j yg berada pada index ke 4

}
