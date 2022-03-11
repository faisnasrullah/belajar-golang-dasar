package main

import "fmt"

func main() {
	/*
		Operasi Perbandingan :
		>	(lebih dari)
		<	(kurang dari)
		>=	(lebih dari atau sama dengan)
		<=	(kurang dari atau sama dengan)
		==	(sama dengan)
		!=	(tidak sama dengan)
	*/

	a := 5
	b := 10
	var result bool = a > b

	fmt.Println(result)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}
