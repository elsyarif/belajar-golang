package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer representasi suatu kejadian
// ketika waktu timer sudah expire, maka event akan di kirim ke dalam channel
// untuk membuat timer bisa menggunakan time.NewTimer(duration)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

// kadang butuh channel nya saja, tidak mebutuhkan data timernya
// untuk hal ini bisa menggunakan function time.After(duration)
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

// kadang ada kebutuhan ingin menjalankan sebuah function dengan delay waktu tertentu
// bisa memanfaatkan Timer dengan menggunakan function time.AfeterFunc()
// tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan di panggil 	ketika timer mengirim kejadian
func TestAfeterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Dijalankan setelah 1 detik:")
		group.Done()
	})

	group.Wait()
}
