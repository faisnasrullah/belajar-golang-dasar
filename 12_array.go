package main

import "fmt"

func main() {
	/*
		Tipe Data Array :
		- kumpulan data dengan tipe yg sama
		- tentukan jumlah data yg bisa ditampung oleh array tersebut
		- daya tampung tidak bisa bertambah ketika array sudah dibuat

		Function Array :
		- len(array)				=> untuk mendapatkan panjang array
		- array[index]				=> mendapatkan data pada index yg ditentukan
		- array[index] = value		=> mengubah data pada index yg ditentukan
	*/

	// cara pertama
	var friends [3]string // hanya bisa menampung 3 data saja
	friends[0] = "Eko Sungkowo"
	friends[1] = "Fais Nasrullah"
	friends[2] = "Farid Afriyanto"

	fmt.Println(friends[1])

	// cara kedua
	var skills = [4]string{
		"JavaScript",
		"Python",
		"Go",
		"Dart",
	}

	fmt.Println(skills)
	fmt.Println(len(skills))
}
