package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	/**
	Nil :
	- saat kita membuat variabel dengan tipe data seperti (interface, function, map, slice, pointer dan channel)
	  makan secara otomatis akan dibuatkan default valuenya adalah nil.
	- Nil bisa kita gunakan untuk melakukan pengecekan
	*/

	person := NewMap("Jhon Wick")

	// melakukan pengecekan menggunakan nil
	if person == nil {
		fmt.Println("Data Masih Kosong Cuy")
	} else {
		fmt.Println(person)
	}
}
