package main

import "fmt"

/** Nil
* 	data nil yaitu tipe data kosong
* 	Nil sendiri hanya bisa digunkan di beberapa tipe data, seperti :
	- interface,
	- function,
	- map,
	- slice,
	- pointer,
	- channel
**/

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	fmt.Println("=== Nil ===")
	data := newMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
