package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

func (person Person) sayHello() {
	fmt.Println("Hello, My Name is", person.Name)
}

func main() {
	/**
	Struct Method :
	- struct adalah tipe data seperti tipe data yg lainnya, bisa digunakan sebagai parameter untuk function
	- namun jika ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struc memiliki function
	- method adalah function
	*/

	jhon := Person{Name: "Jhon Wick"}
	jhon.sayHello()
}
