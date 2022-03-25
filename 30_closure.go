package main

import "fmt"

func main() {
	/*
		Closures :
		- adalah kemampuan sebuah function berinteraksi dengan data - data disekitarnya dalam scope yg sama
		- harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
	*/

	name := "Fais Nasrullah" // variabel name ini bisa diakses dan diubah dari function increment
	counter := 0             // variabel counter ini bisa diakses dan diubah dari function increment

	increment := func() {
		name = "Ini Budi" // harap berhati hati ketika menggunakan fitur closure
		fmt.Println("Increment")
		counter++ // mengubah variabel counter dari dalam function
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
