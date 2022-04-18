package main

import (
	"fmt"
	"os"
)

/**
	Package OS
	- berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
	- https://golang.org/pkg/os/
	- pastikan mengambil data argument dimulai dari index ke 1, karena index ke 0 adalah hasil kompilasi file tersebut
**/
func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args[0]) // hasil kompilasi
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error : ", err.Error())
	}

	// export TEST_USERNAME=ampasdev
	// export TEST_PASSWORD=123456007
	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")

	fmt.Println("Username : ", username)
	fmt.Println("Password : ", password)
}
