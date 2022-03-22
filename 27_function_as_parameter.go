package main

import "fmt"

func spamFilter(name string) string {
	if name == "Anjing" {
		return "~SENSOR~"
	} else {
		return name
	}
}

// TODO: Type Declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hai", nameFiltered)
}

func main() {
	/**
	Function as Parameter :
	- function tidak hanya bisa kita simpan didalam variabel sebagai variabel
	- function juga bisa kita gunakan sebagai parameter untuk function lain
	*/
	sayHelloWithFilter("Jhon Wick", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
