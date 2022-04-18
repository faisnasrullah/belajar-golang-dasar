package main

import (
	"flag"
	"fmt"
)

/**
	Package Flag
	- berisikan fungsionalitas untuk memparsing command line argument
	- https://golang.org/pkg/flag/
**/

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host") // name, default value, description
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "toor", "Put your database password")
	var number *int = flag.Int("number", 88, "Put your favorite number")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)
}
