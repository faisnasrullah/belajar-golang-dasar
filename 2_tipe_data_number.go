package main

import "fmt"

func main() {
	/*
		Tipe Data Number :
		1. Integer
		  - int8	[min: -128, max: 127]
		  - int16	[min: -32768, max: 32767]
		  - int32	[min: -2147483648, max: 2147483647]
		  - int64	[min: -9223372036854775808, max: 9223372036854775807]
		  - uint8	[min: 0, max:255]
		  - uint16	[min: 0, max:65535]
		  - uint32	[min: 0, max:4294967295]
		  - uint64	[min: 0, max:18446744073709551615]

		2. Floating Point
		  - float32
		  - float64
		  - complex64
		  - complex128

		Alias :
		  - byte == uint8
		  - rune == int32
		  - int == minimal int32 (depend on your OS 32bit or 64bit)
		  - uint == minmal uint32 (depend on your OS 32bit or 64bit)
	*/

	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma Lima = ", 3.5)
}
