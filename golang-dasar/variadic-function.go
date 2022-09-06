package main

import (
	"fmt"
)

// parameter yang berada di posisi akhir, memiliki kemampuan dijadikan sebuah varargs
// varargs artinya datanya bisa menerima lebih dari satu input, atau angap saja sebagai array
// perbedaan parameter biasa dengan tipe data array
// - array : wajib memebuat array terlebih dahulu sebelum mengirim ke function
// - varargs : bisa langsung mengirim datanya, jika lebih dari satu cukup gunakan tanda koma (,)

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Slice parameter
// varirable berupa slice
// bisa menjadikan slice sebagai vararg parameter

func main() {
	total := sum(10, 11, 9, 43, 9, 33)

	fmt.Println(total)

	fmt.Println("--=== Slice Parameter ===--")
	numbers := []int{10, 10, 10, 10, 10}
	total = sum(numbers...) // slice menjadi parameter
	fmt.Println(total)
}
