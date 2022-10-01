package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
}

func CreateCounter(ctx context.Context) chan int {
	distination := make(chan int)

	go func() {
		defer close(distination)

		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				distination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return distination
}

//func TestContextWithCancel(t *testing.T) {
//	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
//
//	distination := CreateCounter()
//
//	for n := range distination {
//		fmt.Println("Counter", n)
//		if n == 10 {
//			break
//		}
//	}
//
//	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
//}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	distination := CreateCounter(ctx)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())

	for n := range distination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke contex
	time.Sleep(3 * time.Second)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	distination := CreateCounter(ctx)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())

	for n := range distination {
		fmt.Println("Counter", n)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	distination := CreateCounter(ctx)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())

	for n := range distination {
		fmt.Println("Counter", n)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Totoal Goroutine", runtime.NumGoroutine())
}
