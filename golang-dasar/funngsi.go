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
	fmt.Println("Hello,", name, lastname)
}

// function return value
// function bisa mengembalikan data
// untuk mengembalikan data, kita harus menuliskan tipe data
// untuk mengembalikan data dari fungsi, kita bisa menggunakan kata kunci return di ikuti dengan datanya
func witHello(param string) string {
	return "Hello, " + param
}

// funcion multiple value
// untuk memberitahu jika fungsi mengembalikan multiple value kita harus menuliskan semua tipe data return value di gunski
// jika kita ingin menhiraukan retun value tersebut, kita bisa menggunakan tanda _ (garis bawah)
func getFullName() (string, string) {
	return "Syarif", "Hidayatulloh"
}

// named return value
// biasanya kita memberi tahu bahwa sebuah function mengembalikan value, maka kita kita hanya mendeklerasikan tipe data return di function
// kita bisa membuat variable secara langsung di tipe data return function nya
func getColmplateName() (firstname, lastname, nickname string) {
	firstname = "Syarif"
	lastname = "Hidayatulloh"
	nickname = "arif"

	return "Hi, nama saya ", firstname + " " + lastname, "Bisa di panggil " + nickname
}

func main() {
	sayHello()
	sayHelloTo("syarif", "Hidayatulloh")
	param := witHello("param")
	fmt.Println(param)
	fmt.Println("--==== return multiple value ====--")
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println("--==== return multiple value ====--")
	firstName1, _ := getFullName() // _ untuk menghiraukan return lastName
	fmt.Println(firstName1)
	fmt.Println("--==== return named value ====--")
	namaAwal, namaAkhir, pangilan := getColmplateName()
	fmt.Println(namaAwal, namaAkhir, pangilan)
}
