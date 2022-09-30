package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// Map ini mirip dengan Go-lang Map, namun yang membedakan ini aman untuk menggunakan
// concurent menggunakan goroutine
// fungsi yang bisa gigunakan di Map
//	Stroe(key, value) untuk menyimpan data ke map
//	Load(key) untuk mengambil data dari Map menggunakan Key
//	Delete(key) untuk menghapus data di Map menggunakan key
//	Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
