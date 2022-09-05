package main

import (
	"fmt"
)

func main() {
	// slice adalah potongan data dari array
	// slice mirip dengan array,yang membedakan ukuran slice bisa berubah
	// slice adalah data yang mengakses sebagian atau seluruh data array
	// slice memiliki 3 data : Pointer,length dan capatity
	// Pointer adalah penunjuk data pertama di array para slice
	// Length adalah panjang dari slice, dan
	// Capatity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	/* Membuat slice dari array */
	// array[low:high] //membuat slice dari array dimulai dari index low sampai index sebelum high
	// array[low:] //membuat slice dari array dimulai index low sampai index akhir di array
	// array[:high] //membuat slice dari array dimulai index 0 sampai index sebelum high
	// array[:] //membuat slice dari array dimulai index 0 sampai index akhir di array
	names := [...]string{"Syarif", "Hidayatulloh", "AL", "Haidar", "Ahmad"}
	slice := names[2:4] // membuat slice dari array di mulai dari index 2 sampai index sebelum 4

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	/** Fungsi pada slice
	len(slice)			//untuk mendapatkan panjang
	cap(slice) 			//untuk mendapatkan kapasitas
	append(slice, data) //membuat slice baru dengan menambahkan data ke posisi terakhir slic, jika kapasitas sudah penuh, maka akan membuat array baru
	make([]TypeData, length, capacity)
	*/

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"} // array days
	fmt.Println("days before :", days)
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu new"
	daySlice1[1] = "Minggu new"
	// hati-hati saat mengubah slice, karena akan merubah nilai array nya juga
	fmt.Println("=================================")
	fmt.Println("days :", days)

	daySlice2 := append(daySlice1, "Libur baru") // menambahkan data baru di posisi terakhir
	daySlice2[0] = "Ups"                         // menambahkan data pada index ke 0
	fmt.Println(daySlice2)
	fmt.Println(days)

	// kode membuat slice dengan menggukanan fungsi make
	s := make([]string, 4)

	s[0] = "a"
	s[1] = "b"
	fmt.Println("slice :", s)
	fmt.Println("Len slice :", len(s))
	fmt.Println("Cap slice :", cap(s))

	//copy slice dari slice yang lain
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // membuat slice dengan panjang dari fromSLice dan capasitas fromSlice

	copy(toSlice, fromSlice)
	fmt.Println("copy slice: ", toSlice)
	// Note: perhatikan saat membuat array, jika salah mebuatnya bukan array yang dibuat melainkan slice
	// perbedaan array dan slice
	isArray := [...]int{1, 2, 4, 5, 6}
	isSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(isArray)
	fmt.Println(isSlice)
}
