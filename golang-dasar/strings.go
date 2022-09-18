package main

import (
	"fmt"
	"strings"
)

// Package stingsa adalah package yang berisiakan function-function untuk memanipulasi tipe data string

func main() {
	fmt.Println(strings.Contains("Syarif Hidayatulloh", "Syarif")) // untuk mengecek apakah s: mengandung string lain
	fmt.Println(strings.Trim("syarif hi   ", " "))                 // untuk memotong cutset di awal dan di akhir
	fmt.Println(strings.ToLower("SYARIF HIDAYATULLOH"))            // untuk membuat semua karakter sting menjadi lower case
	fmt.Println(strings.ToUpper("syarif Hidayatulloh"))            // untuk membuat semua karakter string menjadi upper case
	fmt.Println(strings.Split("Syarif Hidayatulloh", " "))         // untuk memotong string berdasarkan sparator
	fmt.Println(strings.ReplaceAll("Syrif", "Syrif", "Syarif "))   // untuk mengubah string yang lama ke yang baru
}
