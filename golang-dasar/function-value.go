package main

import (
	"fmt"
)

// Function adalah first class citizen
// function juga merupaka tipe data, dan bisa disimpan di dalam variable

func getGoodBay(name string) string {
	return "Good by, " + name
}

func main() {
	goodbay := getGoodBay          // function di tampung dalam variable
	fmt.Println(goodbay("Syarif")) // variable di panggil denga memasukan parameter
}
