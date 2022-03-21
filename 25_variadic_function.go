package main

import "fmt"

// TODO: Basic penggunaan variadic function
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	/**
	Variadic Function :
	- parameter yg berada diposisi paling akhir, memiliki kemampuan dijadikan sebuah varargs
	- varargs adalah datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
	- apa bedanya dengan parameter biasa dengan tipe data array? :
		@ jika parameter tipe array -> kita wajib membuat array terlebih dahulu sebelum mengkirimkan ke function
		@ jika parameter menggunakan varargs -> kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma
	Ciri ciri variadic function :
	- terdapat tiga tanda titik sebelum menuliskan jenis tipe data, contoh : func sumAll(numbers ...int) int {}
	- varargs hanya bisa satu dan berada dipaling akhir sebagai parameter
	*/

	total := sumAll(10, 5, 4, 23, 70, 33)

	// menggunakan slice sebagai variabel argument
	angka := []int{32, 55, 10, 2, 8, 99}
	hasil := sumAll(angka...)

	fmt.Printf("Totalnya adalah %v\n", total)
	fmt.Printf("Jumlahnya adalah %v\n", hasil)
}
