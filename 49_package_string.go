package main

import (
	"fmt"
	"strings"
)

/**
	Package String
	strings.Trim(string, cutset)		=>		memotong cutset diawal dan akhir string
	strings.ToLower(string)				=>		membuat semua karakter string menjadi lower case
	strings.ToUpper(string)				=>		membuat semua karakter string menjadi upper case
	strings.Split(string, separator)	=>		memotong string berdasarkan separator
	strings.Contains(string, search)	=>		mengecek apakah string mengandung string lain
	strings.Replace(string, from, to)	=>		mengubah semua string dari from ke to
**/

func main() {
	country := "Indonesia"
	poem := "Negara Kesatuan Republik Indonesia"
	sliceOfPoem := strings.Split(poem, " ")
	fmt.Println("Slice Of Poem, index ke 2 :", sliceOfPoem[2])

	fmt.Println(strings.Trim("  fais nasrullah    ", " ")) // menghilangkan spasi yg ada di sebelah kiri dan kanan
	fmt.Println(strings.ToLower(country))
	fmt.Println(strings.ToUpper(country))
	fmt.Println((strings.Split(poem, " "))) // hasilnya akan menjadi slice of string
	fmt.Println(strings.Contains(country, "Indo"))
	fmt.Println(strings.Replace("Hellow World", "World", "Jhon Wick", 1)) // me replace kata yg sesuai dengan index nya
	fmt.Println(strings.ReplaceAll("Hello hello Hello", "hello", "haii")) // me replace semua kata yg cocok
}
