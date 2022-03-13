package main

import "fmt"

func main() {
	/*
		Tipe Data Slice :
		- adalah potongan dari tipe data array
		- slice mirip dengan array, yg membedakan ukuran slice bisa berubah
		- slice sebenarnya menyimpan data pada array
		- slice memiliki 3 data (pointer, length and capacity)
		- pointer 	-> petunjuk data pertama pada array of slice
		- length	-> panjang dari slice tersebut
		- capacity	-> kapasitas dari slice, dimana length tidak boleh lebih dari capacity

		Membuat Slice dari Array :
		- array[low:high]		=>	Membuat slice dari array dimulai dari index low sampai index sebelum high
		- array[low:]			=>	Membuat slice dari array dimulai dari index low sampai index terakhir
		- array[:high]			=>	Membuat slice dari array dimulai dari index ke 0 sampai index sebelum high
		- array[:]				=>	Membuat slice dari array dimulai dari index ke 0 sampai index terakhir

		Function pada Slice :
		- len(slice)							=> Untuk mendapatkan panjangdf dari sebuah slice
		- cap(slice)							=> Untuk mendapatkan kapasitas dari sebuah slice
		- append(slice, data)					=> Menambahkan data ke posisi terakhir, jika capacity sudah penuh, akan membuat array baru
		- make([]TypeData, length, capacity)	=> Membuat slice baru
		- copy(destination, source)				=> Menyalin slice dari source ke destination
	*/

	// membuat array terlebih dahulu
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// membuat slice dari array of months
	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// saat kita mengubah data array, slice akan secara otomatis mengikuti value dari si array tersebut
	months[5] = "Juni Nih"
	fmt.Println(months)
	fmt.Println(slice1)

	// begitupun sebaliknya, jika kita mengubah value dari slice yg kita buat dari array, data pada array akan ikut terubah
	slice1[0] = "Mei ya ini"
	fmt.Println(months)
	fmt.Println(slice1)

	slice2 := months[10:]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := append(slice2, "Ramadhan") // dikarenakan capacity dari slice2 adalah 2, maka akan otomatis membuat array baru
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println(slice2)
	fmt.Println(months)

	// membuat slice baru dengan function make()
	days := make([]string, 2, 6)
	days[0] = "Senin"
	days[1] = "Selasa"

	fmt.Println(days)
	fmt.Println(len(days))
	fmt.Println(cap(days))

	// menyalin slice ke slice yg lain dengan function copy()
	copyDays := make([]string, len(days), cap(days))
	copy(copyDays, days)
	fmt.Println(copyDays)

	// hati hati saat membuat array (perbedaan array dengan slice)
	iniArray := [3]string{
		"andi",
		"budi",
		"cici",
	}
	iniJugaArray := [...]int{1, 2, 3}
	iniSlice := []int{3, 2, 1}

	fmt.Println(iniArray)
	fmt.Println(iniJugaArray)
	fmt.Println(iniSlice)
}
