package main

import (
	"fmt"
	"go-lang-dasar/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
