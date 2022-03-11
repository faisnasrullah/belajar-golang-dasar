package main

import "fmt"

func main() {
	/*
		Operasi Boolean :
		&&	(AND)
		||	(OR)
		!	(XOR / Kebalikan)

		&&	=> hanya bernilai true jika keduanya bernilai true
		||	=> hanya bernilai false jika keduanya bernilai false
		!	=> jika bernilai false, maka hasilnya akan menjadi true
	*/

	var nilaiAkhir = 90
	var nilaiAbsensi = 70

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiAbsensi = nilaiAbsensi > 70
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusNilaiAbsensi)

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi

	fmt.Println(lulus)
}
