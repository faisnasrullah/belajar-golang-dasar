package main

import "fmt"

func main() {
	/*
		Konversi Tipe Data :

	*/
	title := "Belajar Go Lang Dasar"
	var e byte = title[1] // jika kita mengambil satu karakter == byte
	fmt.Println(string(e))

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
}
