package main

import "fmt"

// tanpa menggunakan recursive function
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// solusi menggunakan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	/**
	Recursive Function :
	- Function yg memanggil dirinya sendiri
	- contoh kasus yg menggunakan recursive function adalah factorial
	*/

	loop := factorialLoop(5)

	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)

}
