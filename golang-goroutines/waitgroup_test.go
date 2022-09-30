package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// waitgroup adalah sebuah fiture yang bisa gigunakan untuk menunggu sebuah proses selesai dilakukan
// hall ini kadang diperlukan: ingin menjalankan beberapa proses dengan goroutine tapi kita ingin
// semua proses selesai terlebih dahulu sebelum aplikasi selesai dijalankan
// untuk menandai bahwa ada proses goroutine, bisa menggunakan method Add(int),
//	setelah proses goroutine selesai bisa menggunakan Done()
// 	untuk menugnggu semua proses selesai, bisa menggunakan method Wait()

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Println("Completed")
}
