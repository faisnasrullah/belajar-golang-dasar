package main

import "fmt"

func endApp() {
	fmt.Println("End App")
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
	Panic :
	- adalah function yang bisa kita gunakan untuk menghentikan program
	- biasanya dipanggil ketika terjadi error pada saat program kita berjalan
	- saat panic function dipanggil, program akan berhenti, namun defer function tetap akan dieksekusi
	*/

	runApp(true)
}
