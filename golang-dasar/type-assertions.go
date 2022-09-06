package main

import "fmt"

func random() interface{} {
	return "true"
}

// assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// assertions sering kali digunakan ketika bertemu dengan data interface kosong

func main() {
	result := random()
	resulString := result.(string)
	fmt.Println(resulString) // panic

	// saat salah menggunakan type assertions, maka akan berakibat panic di aplikasi
	// jika panic tidak tercover, maka otomatis program akan terhenti
	// agar lebih aman,sebaiknya menggunakan switch expression untuk melakukan type assertions

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknow type")
	}
}
