package main

import (
	"fmt"
)

func main() {
	// TODO: Basic for loop
	counter := 1 // init statement

	for counter <= 10 {
		fmt.Printf("Perulangan ke - %v\n", counter)
		counter++ // post statement
	}

	// TODO: Kode program for dengan statement
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke : ", i)
	}

	// TODO: For Range
	// ? digunakan untuk melakukan iterasi terhadap semua data collection (array, map, slice)
	friends := []string{"Eko", "Fais", "Farid", "Hasan"} // type data slice

	for index, value := range friends {
		fmt.Printf("Index : %v - Value : %v\n", index, value)
	}

	person := make(map[string]string)
	person["name"] = "Fais Nasrullah"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
