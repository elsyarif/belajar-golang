package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func getAverage(numbers []int, ch chan float64) {
	sum := 0

	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	ch <- max
}

func TestMain(t *testing.T) {
	runtime.GOMAXPROCS(3)

	numbers := []int{3, 4, 3, 5, 99, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println(numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("max \t: %d \n", max)
		}
	}
}
