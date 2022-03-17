package main

import "fmt"

func main() {
	name := "Fais Nasrullah"

	if name == "Fais" {
		fmt.Printf("Hai %v, selamat belajar if expression pada go\n", name)
	} else if name == "Nasrullah" {
		fmt.Println("Hai selamat belajar golang")
	} else {
		fmt.Println("Hai, kenalan dong!")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang bre")
	} else {
		fmt.Println("Nah gini dong mudah diingat namanya")
	}
}
