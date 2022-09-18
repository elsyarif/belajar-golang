package helper

import "fmt"

// Access Modifier
var version = "1.0.0"              // tidak bisa di akese dari luar
var Application = "Belejar Golang" // Datpat di akses dari luar

func SayHello() {
	fmt.Println("Hello")
}

func sayGoodbay() {
	fmt.Println("Goodbay")
}
