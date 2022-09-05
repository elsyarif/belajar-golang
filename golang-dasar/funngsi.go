package main

import "fmt"

// funsi merukan sebuah kode blok yang sengaja dibuat dalam perogram agar bisa digunakan berulang-ulang
// membuat fuction, hanya menggunakan kata kunci func diikuti dangan nama funsi dan diikuti dengan kode blok
// memanggil function dengan kata kunci nama func di ikuta kurung buka kurung tutup
func sayHello() {
	fmt.Println("Hello")
}

// Function dengan parameter
// ketika parameter pada function, maka kita wajib memasukan data ke parameternya
func sayHelloTo(name string, lastname string) {
	fmt.Println("Hello, ", name, " ", lastname)
}

// function return value
// function bisa mengembalikan data
// untuk mengembalikan data, kita harus menuliskan tipe

func main() {
	sayHello()
	sayHelloTo("syarif", "Hidayatulloh")
}
