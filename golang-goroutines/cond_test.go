package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond adalah locking berbasis kondisi
// cond membutuhkan locker(bisa menggunakan Mutex atau RWMutex) untuk implementasi locking
// Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
// Function Signal() bisa digunakan untuk memberitahu sebuah goroutine agar tidak perlu menunggu lagi
// Function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done", value)
	cond.L.Unlock()
	group.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)

		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()
}
