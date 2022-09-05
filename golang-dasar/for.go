package main

import "fmt"

func main() {
	// for digunkan biasanya untuk perulangan

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	// for bisa digunakan untuk melakukan data iterasi semua data collection
	// data collection berupa, Array, slice, map
	names := []string{"Syarif", "Hidayatulloh", "Dayat", "Hida"}
	for index, name := range names {
		fmt.Println("Index ke-", index, " = ", name)
	}

	// Breake & Continue adalah kata kunci yang biasa digunakan dalam perulangan
	// breake digunakan untuk menghentikan seluruh perulangan
	// continue digunakan untuk perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya
	fmt.Println("Contoh perulangan dengan break")
	for i := 0; i < 10; i++ {
		if i == 7 {
			break // akan menghentikan perulangan.
		}

		fmt.Println("Perulangan ke-", i)
	}

	fmt.Println("Contoh perulangan dengan continue")
	for i := 0; i <= 6; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke-", i)
	}
}
