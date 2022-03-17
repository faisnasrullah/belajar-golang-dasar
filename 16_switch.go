package main

import "fmt"

func main() {
	hari := "Kamis"

	switch hari {
	case "Senin":
		fmt.Println("Hari Senin")
	case "Selasa":
		fmt.Println("Hari Selasa")
	case "Rabu":
		fmt.Println("Hari Rabu")
	case "Kamis":
		fmt.Println("Hari Kamis")
	case "Ju'mat":
		fmt.Println("Hari Ju'mat")
	case "Sabtu":
		fmt.Println("Hari Sabtu")
	case "Minggu":
		fmt.Println("Hari Minggu")
	default:
		{
			fmt.Println("Haa?")
			fmt.Println("Bukan Nama Hari Bre Ini...")
		}
	}

	// switch short statement
	switch length := len(hari); length >= 5 {
	case true:
		fmt.Println("Nama hari lebih dari atau sama dengan 5 karakter")
	case false:
		fmt.Println("Nama hari kurang dari atau sama dengan 5 karakter")
	}

	// switch tanpa kondisi (lebih mirip seperti if)
	name := "Fais N"
	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Mudah di Ingat")
	}
}
