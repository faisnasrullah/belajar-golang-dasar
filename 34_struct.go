package main

import "fmt"

// contoh struct customer
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	/**
	Struct :
	- sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	- biasanya merepresentasikan data dalam program aplikasi yg kita buat
	- data struct disimpan dalam field
	- sederhanannya struct adalah kumpulan dari field
	- struct bisa dibilang mirip class jika pada bahasa pemrogramman berorientasi objek
	*/

	// cara membuat data struct
	// cara 1 :
	var customer1 Customer
	customer1.Name = "Eko Sungkowo"
	customer1.Address = "Tangerang"
	customer1.Age = 30

	fmt.Println(customer1)

	// struct literal
	// sebenarnya banyak cara yang bisa kita gunakan untuk membuat data dari struct
	// cara 1 :
	// penulisan ini lebih disarankan, karena urutan penulisannya boleh bolak balik atau tidak berurutan
	customer2 := Customer{
		Name:    "Farid Afriyanto",
		Address: "Pekalongan",
		Age:     25,
	}

	fmt.Println(customer2)

	// cara 2 :
	// cara ini lebih singkat, namun tidak direkomendasikan, karena jika ada urutan ataupun penambahan pada struct, maka akan terjadi error
	customer3 := Customer{"Budi", "Bandung", 22}
	fmt.Println(customer3)

}
