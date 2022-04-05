package database

import "fmt"

var connection string

/**
Function Init :
- sebuah function yg diakses secara otomatis ketika package diakses

Blank Identifier :
- maksudnya adalah jika kita melakukan import sebuah package,
  namum belum menggunakan function atau variabel yg ada di package tersebut
  maka golang akan secara otomatis menghapus import package tersebut
- untuk mengatasi hal itu, kita bisa menggunakan _ sebelum nama package, contoh :
import _ "fmt"
*/
func init() {
	fmt.Println("Function Init Running")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
