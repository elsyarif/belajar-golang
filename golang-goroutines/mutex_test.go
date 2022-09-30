package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// untuk mengatasi malsah race condition, dengan struct sync.Mutex
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

// kadang ada kasus dimana ingin melakukan locking tidak hanya pada proses mengubah data,
// tapi juga membaca data, untuk melakukannya menggunakan RWMutex memiliki dua jenis lock
// lock untuk Read, dan lock untuk Write
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
// hati hati saat membuat aplikasi yang parallel atau councurent, masalah yang sering kita hadapi adalah deadlock
// deadlock dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa berjalan

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name: "Syarif",
	}

	user2 := UserBalance{
		Name: "Hidayatulloh",
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(5 * time.Second)

}
