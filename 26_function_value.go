package main

import "fmt"

func sayGoogBye(name string) string {
	return "Good Bye " + name
}

func main() {
	/**
	Function Value :
	- function adalah first class citizen
	- function juga merupakan tipe data, dan bisa kita simpan didalam variable,
	*/

	//? Normalnya kita menggunakan function
	fmt.Println(sayGoogBye("Fais"))

	//? Namun kita bisa menyimpan function kedalam variabel
	goodBye := sayGoogBye
	fmt.Println(goodBye("Jhon Wick"))
}
