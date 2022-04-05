package main

import (
	"fmt"
	"go-lang-dasar/helper"
)

/**
Package :
- mengakses package helper dari file lain kita harus menggunakan import

Access Modifier :
- kita sudah buat beberapa case di package helper
- variabel yg bisa diakases -> Application
- variabel yg tidak bisa diakases -> version
- function yg bisa diakases -> SayHello
- function yg tidak bisa diakases -> sayGoodbye
*/

func main() {
	result := helper.SayHello("Jhon Wick")
	fmt.Println(result)
	fmt.Println(helper.Application)
	// ACCESS MODIFIER (yang tidak bisa diakses)
	// goodbye := helper.sayGoodBye("Jhon") // error karena tidak bisa diakses
	// fmt.Println(helper.version) // error karena tidak bisa diakses
}
