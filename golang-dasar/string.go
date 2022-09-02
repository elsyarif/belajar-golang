package main

import "fmt"

func main() {
	fmt.Println("Syarif")
	fmt.Println("Syarif Hidayatulloh")

	var nama string
	nama = "Syarif"
	namalengkap := "Syarif Hidayatulloh"
	fmt.Println(nama)
	fmt.Println(namalengkap)

	// menghitung jumlah karakter string
	fmt.Println(len(nama))

	// Mengambil karakter pada posisi yang ditentukan
	fmt.Println(namalengkap[2])
}
