package main

import "fmt"

func main() {
	/*
		Operasi Matematika :
		+ (penjumlahan)
		- (pengurangan)
		* (perkalian)
		/ (pembagian)
		% (sisa pembabagian)

		Augmented Assignments :
		a += 10		=>	a = a + 10
		a -= 10		=>	a = a - 10
		a *= 10		=>	a = a * 10
		a /= 10		=>	a = a / 10
		a %= 10		=>	a = a % 10

		Unary Operator :
		a++			=> a = a + 1
		a--			=> a = a - 1
		+			=> Positive (default)
		-			=> Negative
		!			=> Boolean Kebalikan
	*/

	a := 5
	b := 8
	c := a + b

	var i = 70
	i += 20

	fmt.Println(c)
	fmt.Println(i)

	i++
	fmt.Println(i)
}
