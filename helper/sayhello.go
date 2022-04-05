package helper

/**
Package :
- tempat yg digunakan untuk mengorganisir kode program
- package sendiri sebenarnya hanya sebuah folder
- dalam satu folder, tidak bisa menggunakan nama package yg sama
*/

// ACCESS MODIFIER :
// penggunaan nama variabel, jika diawali dengan HURUF BESAR, maka bisa diakses di package / file lain
// begitupun sebaliknya, jika diawali dengan HURUF KECIL, maka tidak bisa diakses di package / file lain
var version = 1
var Application = "Belajar Golang Dasar"

// ACCESS MODIFIER :
// karena nama function diawali dengan HURUF BESAR, maka bisa diakses di package / file lain
func SayHello(name string) string {
	return "Hellow " + name
}

// ACCESS MODIFIER :
// karena nama function diawali dengan HURUF KECIL, maka tidak bisa diakses di package / file lain
func sayGoodBye(name string) string {
	return "Goodbye " + name
}
