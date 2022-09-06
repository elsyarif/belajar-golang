package main

import (
	"fmt"
)

/** Pointer di Function
*	saat membuat function parameter di function, secara default adalah pass by value, artinya
	Data akan copy lalu akan di kirim ke function tersebut
*	Jika mengubah data di dalam function data yang aslinya tidak akan berubah
* 	hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
*	Jika ingin mengubah data asli parameter tersebut, bisa menggunakan pointer function
*	untuk menjadikan sebuah parameter sebagai pointer, bisa menggunakan operator * di parameternya
**/

type Address1 struct {
	City     string
	Province string
	Country  string
}

func changeCounry(address *Address1) {
	address.Country = "Indonesia"
}

func main() {
	address := Address1{
		City:     "Cilegon",
		Province: "Banten",
		Country:  "",
	}

	changeCounry(&address)
	fmt.Println(address)
}
