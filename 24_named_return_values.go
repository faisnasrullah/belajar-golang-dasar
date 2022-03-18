package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Jhon"
	middleName = "F"
	lastName = "Wick"

	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
