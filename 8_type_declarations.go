package main

import "fmt"

func main() {
	/*
		Type Declarations :
		- kemampuan membuat ulang tipe data baru dari tipe data yg sudah ada
		- biasanya digunakan untuk membuat alias terhadap tipe data yg sudah ada
		- contoh type alias bawaan go adalah : byte alias dari tipe data uint8
	*/

	type NoKTP string
	type Married bool

	var noKtpFais NoKTP = "33123131231"
	var marriedStatus Married = true

	fmt.Println(noKtpFais)
	fmt.Println(marriedStatus)
}
