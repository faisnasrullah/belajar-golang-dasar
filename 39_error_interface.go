package main

import (
	"errors"
	"fmt"
)

// sebuah function yg me return int dan interface error
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil // karena error adalah sebuah interface, maka kita bisa menggunakan nil untuk default valuenya
	}
}

func main() {
	/**
	Error Interface :
	- golang memiliki interface ygn digunakan sebagai kontrak untuk error, nama interfacenya adalah error
	-
	*/

	data, err := Pembagi(50, 0)
	if err == nil {
		fmt.Println("Hasilnya :", data)
	} else {
		fmt.Println("Errornya :", err)
	}
}
