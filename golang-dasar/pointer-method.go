package main

import "fmt"

/** Pointer di Method
* 	walaupun method akan menempel di struc, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
*	Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasikan ketika memenggil method
**/

type Man struct {
	Name string
}

// Merried struct menambahkan * untuk menjadika pointe pada method
func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	syarif := Man{Name: "Syarif"}
	syarif.Merried()

	fmt.Println(syarif)
}
