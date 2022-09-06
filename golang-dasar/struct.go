package main

import "fmt"

/**
*	Struct adalah sebuah template data yang bisa digunakan untuk menggabungkan nol atau lebih
* 	tipe data lainny dalam satu kesatuan
* 	Struct biasanya representasi data dalam program yang dibuat
* 	Data struct di simpan dalam field, struct adalah sekumpulan dari field
**/

type Customer struct {
	Name, Address string
	Age           int
}

/**	Membuat data struct
* 	struct adalah template data atau prototype data
* 	Struct tidak bisa langsung digunakan, namun bisa membuat data/objek dari struct yang telah dibuat
**/

/** Struct method
*	bisa digunakan sebagai parameter untuk function
* 	jika kita ingin menambahkan mthod dalam struct, sehingga seakan-akan struct memiliki function
* 	Method adalah function
**/
func (customer Customer) sayHello() {
	fmt.Println("Hello, nama saya ", customer.Name)
}
func main() {
	// Membuat data struct
	var member Customer
	member.Name = "Syarif"
	member.Address = "Cilegon"
	member.Age = 29

	fmt.Println(member)

	// Struct Literal
	joko := Customer{
		Name:    "Joko",
		Address: "Serang",
		Age:     33,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Pandeglang", 22}
	fmt.Println(budi)

	// Struct method
	rully := Customer{Name: "Rully"}
	rully.sayHello()
}
