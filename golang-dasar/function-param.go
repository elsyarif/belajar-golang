package main

import "fmt"

// Function as Parameter
// function tidak hanya disimpan di dalam variable sebagai value
// namun juga bisa digunakan sebagai parameter untuk function lain

// parameter berupa function
func withFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" || name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

// Function type declaration
// jika function terlalu panjang dan agak ribet untuk menuliskan parameter.
// type declaration bisa digunakan untuk membuat alias function, sehingga akan mempermudah menggunakan function sebagai parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello, ", filter(name))
}

func main() {
	withFilter("Syarif", spamFilter)

	filter := spamFilter
	withFilter("Anjing", filter)

	sayHelloWithFilter("Hidayatulloh", filter)
}
