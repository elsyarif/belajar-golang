package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool adalah implementasi desain patern bernama object poll patern
// sederhananya, design patern pool ini digunakan untuk menyimpan data, selanjutnya
// 	untuk menggunakan datanya, bisa mengambil dari pool, dan setelah selesai menggunakan data
// 	bisa menyimpan kembali ke pool nya
// implementasi pool di Go aman dari problem race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Syarif")
	pool.Put("Hidayatulloh")
	pool.Put("AL Ghazzal")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
