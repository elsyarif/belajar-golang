package main

import "fmt"

func main() {
	// constant adalah variable yang nilainy tidak dapat diubah lagi setalah pertama kali diberi nilai
	// Saat pembuataan constant, wajib langsung menginisali datanya
	const namaAwal string = "Syarif " // inisial type data dan langsung diberi nilai nya

	fmt.Println("const: ", namaAwal)

	// declarasi multiple constant
	const (
		fistname string = "Syarif"
		age      int16  = 29
	)

	fmt.Println(fistname)
	fmt.Println(age)
}
