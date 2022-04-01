package main

import "fmt"

func random() interface{} {
	return "200 OK"
}

func main() {
	/**
	Type Assertions :
	- merupakan kemampuan meribah tipe data menjadi tipe data yg diinginkan
	- fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
	*/
	result := random()
	/*
		resultString := result.(string) // kita lakukan type assertion
		fmt.Println(resultString)

			resultInt := result.(int) // akan menghasilkan panic, karena si interface{} dari function string ini mereturn string "OK"
			fmt.Println(resultInt)
	*/

	// kita bisa melakukan type assertion menggunakan switch statement supaya lebih aman
	// dengan menggunakan ini, maka secara otomatis akan sesuai dengan return dari function random() tersebut, apakah string, integer atau lainnya
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknow")
	}
}
