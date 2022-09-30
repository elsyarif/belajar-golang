package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

// sebuah function di package runtime yang bisa kita gunakan untuk mengambil jumlah thread
// atau mengambil jumlah thread
// secara default jumlah thread di Go itu sebanyak jumlah CPU ki komputer kita

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}

func TestChangeGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	// mengubah Thread
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}
