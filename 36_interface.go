package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHellow(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// contoh 1
type Personal struct {
	Name string
}

// contoh 2
type Animal struct {
	Name string
}

// anonymous function yg mengembalikan GetName string (sesuai kontrak pada interface HasName)
// dan menggunakan struct Personal sebagai parameternya
func (personal Personal) GetName() string {
	return personal.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	/**
	Interface :
	- adalah tipe data abstract, dia tidak memiliki implementasi langsung
	- sebuah interface berisikan definisi - definisi method
	- biasanya interface digunakan untuk membuat kontrak dan function yg general
	*/

	var danan Personal
	danan.Name = "Danan Ramadhan"
	SayHellow(danan)

	cat := Animal{
		Name: "Embul",
	}
	SayHellow(cat)
}
