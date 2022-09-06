package main

import "fmt"

/** Interface
*	interface adalah sebuah data abstrak, dia bisa memiliki implementasi langsung
* 	Sebuah interface berisikan definisi-definisi langsung
*	Biasanya interface digunakan sebagai kontrak
**/

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello, ", hasName.GetName)
}

/** Implementasi Interface
*	setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
* 	Sehingga tidak perlu mengimplentasiksa interface secara manual
*	Hal ini agak berbeda dengan habasa pemrograman lain yang ketika membuat interface,
*	kita harus menyebutkan secara eksplisit akan menggunakan interface mana
**/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// type Animal struct {
// 	Name string
// }

// func (animal Animal) GetName() string {
// 	return animal.Name
// }

/*
  - Interface kosong
  - interface kosong adalah interface yang ditak memiliki deklarasi method satupun,
    hal ini membuat secara otomatis semua tipe data yang akan menjadi implementasinya.
  - contoh penggunaa interface kosong
    panic(v interface{})
    recover() interface{}

*
*/

func Ups() interface{} {
	return "Ups"
}

func main() {
	fmt.Pritnln("=== Interface ===")

	person := Person{Name: "Syarif"}
	sayHello(person)

	fmt.Pritnln("=== Interface Kosong ===")
	kosong := Ups()
	fmt.Println(kosong)
}
