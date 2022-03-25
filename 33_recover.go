package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR CUY!!")
	}

	fmt.Println("Aplikasi Berjalan") // tidak dieksekusi ketika error
}

func main() {
	/**
	Recover :
	- adalah function yang bisa kita gunakan untuk menangkap data panic
	- dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
	*/

	runApp(true)
	fmt.Println("Hellow...") // karena panic nya sudah di recover, maka ini akan tetap dieksekusi
}
