package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Printf("Say Hello to %v %v\n", firstName, lastName)
}

func main() {
	sayHelloTo("Fais", "Nasrullah")
}
