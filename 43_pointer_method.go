package main

import "fmt"

type Man struct {
	Name string
}

// gunakan operator * untuk menandakan bahwa method ini menggunakan pointer
// jika tidak menggunakan operator * maka tidak akan terjadi perubahan jika kita panggil dari luar method
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Dari dalam Method :", man.Name)
}

func main() {
	/**
	Pointer di Method :
	- walaupun method akan menempel pada struct, tapi sebenernya data struct yg akan diakses didalam
	  method adalah pass by value
	- sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena
	  harus selalu diduplikasi ketika memanggil method.
	*/

	jhon := Man{"Jhon Wick"}
	jhon.Married()

	fmt.Println("Dari luar method :", jhon.Name)
}
