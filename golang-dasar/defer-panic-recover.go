package main

import (
	"fmt"
)

/**
*	Defer function adalah function yang dijadwalkan untuk dieksekusi
*		setelah sebuah function selesai dieksekusi
*  	Defer akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
**/

func logger() {
	fmt.Println("Selesai mememanggil function")
}

func runApplication() {
	defer logger()
	fmt.Println("Run Application")
}

/**
* 	Panic function adalah function yang bisa di gunakan untuk menghentikan program
* 	Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
* 	saat panic function dipanggil, program akan berhenti, namun defer function tetap akan di eksekusi
**/
func endApplication() {
	fmt.Println("End Application")
	// Penanganan dengan recover setealah program panic
	message := recover()
	fmt.Println("Terjadi error : ", message)
	// setelah penagan panic dengan recover program tetap berjalan normal dengan menampikan pesan error
}

func runApp(error bool) {
	defer endApplication()
	if error {
		panic("Error")
	}
}

/**
*	Recover adalah function yang bisa digunakan untuk menangani data panic
* 	Dengan Recover proses panic akan terhenti, sehingga program akan tetap berjalan
**/

func main() {
	runApplication()
	runApp(true) // saat program berjalan dan terjadi error makan program akan berhenti
}
