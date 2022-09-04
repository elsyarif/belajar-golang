package main

import "fmt"

func main() {
	var nama1 = "syarif"
	var nama2 = "Syarif"

	var result bool = nama1 == nama2

	fmt.Println(result)

	// Operasi Boolean
	var nilaiAkhir = 90
	var Absent = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = Absent > 80
	var lulus bool = lulusAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
}
