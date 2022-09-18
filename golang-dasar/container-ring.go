package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// implementasi struktur data circular list
// circular ring adalah struktur data ring, dimana akhir element akan kembali ke elemen awal
func main() {
	data := ring.New(5) // membuat new circular ring, dengan jumlah data 5

	for i := 0; i < data.Len(); i++ {
		num := strconv.Itoa(i)
		data.Value = "value " + num

		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
