package main

import "fmt"

func main() {
	/** map adalah data lain yang berisikan kumpulan data yang sama,
	* 	namun kita bisa menentukan jenis tipe data index yang akan digunakan.
	*	map kumpulan key-value, dimana key bersifat unik, tidak boleh sama
	* 	Jumlah data yang dimaksukan kedalam map boleh sebanyak-banyak, asalkan kata kunci berbeda,
	* 	jika digunakan key yang sama maka otomatis akan digantikan dengan data yang beru.
	**/
	// untuk membuat map kosong menggunakan fungsi : make(map[type-date]type-data)
	person := map[string]string{
		"name":    "Syarif",
		"address": "cilegon",
		"status":  "menikah",
	}

	fmt.Println("maps :", person)               // menampikan semua data map
	fmt.Println("Name :", person["name"])       // hanya menampikan nama
	fmt.Println("Address :", person["address"]) // hanya menampikan address

	/** Fungsi pada map
	*	len(map)	untuk mendapatakan jumlah data di map
	*	map[key]	mengambil data di map dengan menggunakan key
	* 	map[key] = value 	mengubah data di map dengan value
	*	make(map[type-date]type-data) 	membuat map baru
	* 	delete(map, key)	menghapus data di map dengan key
	 */
	product := make(map[string]any)
	product["title"] = "Laptop MSI"
	product["brand"] = "MSI"
	product["pice"] = 1234

	delete(product, "price")

	fmt.Println(product)
}
