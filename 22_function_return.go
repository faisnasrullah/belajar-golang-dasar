package main

import "fmt"

// contoh function dengan return value bertipe data string
func sayFullname(fullName string) string {
	if fullName == "" {
		return "Hello bre"
	}
	return "Hai " + fullName
}

func main() {
	result := sayFullname("Fais Nasrullah")

	fmt.Println(result)
	fmt.Println(sayFullname(""))
}
