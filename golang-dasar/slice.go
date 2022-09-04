package main

import "fmt"

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
}
