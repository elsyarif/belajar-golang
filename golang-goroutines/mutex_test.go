package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// kapan menggunakan mutex : ketika sebuah variable melakukan sharing yang diakaes beberapa goroutines
// untuk menangani race condition, lakukan locking, dan unlocing
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
}

type BankAcount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAcount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAcount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAcount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Account balance", account.GetBalance())
}

// error deadlock
