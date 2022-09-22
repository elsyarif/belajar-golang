package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number:", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
