package main

import "fmt"

func main() {
	// deklarasi adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	// deklarasi digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar mudah dimengerti
	type NoKTP string
	type Meried bool

	var ktpKu NoKTP = "627829"
	var meried Meried = true

	fmt.Println(ktpKu)
	fmt.Println(NoKTP("89329382"))
	fmt.Println(meried)
}
