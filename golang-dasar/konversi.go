package main

import "fmt"

func main() {
	// konversi dari tipe ke tipe data lain
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32) //konversi ke tipe data int32

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// konversi tipe data string
	var nama = "Syarif hidayatulloh"
	var e byte = nama[0]
	var eString string = string(e)

	fmt.Println(nama)
	fmt.Println(eString)
}
