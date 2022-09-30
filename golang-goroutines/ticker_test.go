package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Ticker prepresentasi kejadian berulang
// ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
// untuk membuat ticker, bisa menggunakan time.NewTicker(duration)
// untuk menghentikan ticker, bisa menggunakan ticker.Stop()

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

// kadang hanya butuh data ticker nya, kadang butuh channelnya
// jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
// mengembalikan ticker, hanya mengembalikan channel timernya saja
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}
