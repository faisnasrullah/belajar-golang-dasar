package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip jika kondisi ini terpenuhi
		}

		fmt.Printf("%v adalah Angka Ganjil\n", i)
	}
}
