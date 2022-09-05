package main

import "fmt"

func main() {
	// array adalah type data yang berisikan tipe data yang sama
	var names [3]string
	names[0] = "syarif"
	names[1] = "Hidayatulloh"

	fmt.Println(names[0])
	fmt.Println(names[1])

	// Membuat array langsung
	var volumes = [3]int{90, 92, 99}

	fmt.Println(volumes)

	// Fungsi array
	fmt.Println(len(volumes)) //untuk mendapatkan panjang array
	fmt.Println(volumes[2])   //untuk mendapatkan data diposisi index
	volumes[1] = 78           //untuk mengubah data diposisi index
	fmt.Println(volumes[1])
}
