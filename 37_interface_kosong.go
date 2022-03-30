package main

import "fmt"

// contoh interface kosong
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups..."
	}
}

func main() {
	/**
	Interface Kosong :
	- adalah interface yang tidak memiliki deklarasi method satupun,
	- hal ini akan membuat secara otomatis semua tipe data akan menjadi implementasinya
	- contoh penggunaan interface kosong pada go :
	1. fmt.Println(a ...interface{})
	2. panic(v interface{})
	3. recover() interface{ }
	*/

	// cara membuat variabel untuk interface kosong
	// var data int = Ups(1) // ini salah
	var data interface{} = Ups(4)
	fmt.Println(data)
}
