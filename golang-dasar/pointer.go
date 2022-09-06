package main

import "fmt"

/** Pointer
*	golang semua variable itu di passing by value, bukan bu reference
*	Jika mengirim sebuah variable kedalam function, method atau variable lain sebenarnya yang dikirim adalah duplikasi value nya
* 	Pinter adalah kemampuan membuat reference ke lokasi data dari di memory yang sama, tanpa meduplikasi data yang sudah ada
* 	dengan kemampuan pointer, bisa membuat pass by reference
*	untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, bisa menggunakan Operator (&)di ikuti dengan nama variable nya
*	&product ==> Pointer
*	saat mengubah variable pointer, maka perubahan hanya variable tersebut
*	semua varible yang mengacu ke data yang sama tidak akan berubah
*	jika ingin merubah seluruh variable yang mengacu ke data tersebut, bisa menggunakan Operator (*)
*	*product ==> pointer
* 	clue:
	(&) variabel hanya berubah pada varible tersebut
	(*) variabel data yang merefer akan berubah
**/

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// Pass by value
	// address1 di simpan ke address2
	address1 := Address{"Cilegon", "Bantern", "Indonesia"}
	//address2 := address1
	address2 := &address1 // menambahkan & , menjadi pointer ke address1, address2 merefer ke address1
	address3 := Address{"Subang", "Bandung", "Indonesia"}
	address2.City = "Serang"

	*address2 = Address{"Jakarta", "DKI", "Indonesia"} // jika menabahkan * variable yang mengacu pada address1 akan berubah

	fmt.Println(address1) // address1 tidak ada perubahan
	fmt.Println(address2)
	fmt.Println(address3) // address 3 tidak berubah karena tidak mengacu pada address1

	// Functio new
	// digunakan untuk membuat pointer, namun function new hanya mengembalikan pointer ke data koson, artinya tidak ada data awal

	address4 := new(Address)
	address4.Country = "Indonesia"
	fmt.Println(address4)
}
