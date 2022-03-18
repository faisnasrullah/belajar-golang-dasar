package main

import "fmt"

// contoh function dengan mengembalikan lebih dari satu data
func person(fullName string, title string) (string, string) {
	return fullName, title
}

func main() {
	fullName, title := person("Fais Nasrullah", "Programmer")
	fmt.Printf("Nama : %v - Profesi : %v\n", fullName, title)
}
