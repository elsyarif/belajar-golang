package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// Once adalah fitur di Go yang bisa digunakan untuk memastikan bahwa sebauh function di eksekusi hanya sekali
// jadi berapapun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusnya

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()

	}

	group.Wait()
	fmt.Println(counter)

}
