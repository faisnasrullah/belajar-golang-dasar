package main

import (
	"fmt"
	"go-lang-dasar/helper"
)

/**
Package :
- mengakses package helper dari file lain kita harus menggunakan import
*/

func main() {
	result := helper.SayHello("Jhon Wick")
	fmt.Println(result)
}
