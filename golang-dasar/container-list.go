package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New() // membuat new container list

	data.PushBack("Syarif")
	data.PushBack("Hidayatulloh")
	data.PushFront("el")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// mengambil semua data dari depan ke belakang
	fmt.Println("=== ")
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println("=== data data belakang ===")
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
