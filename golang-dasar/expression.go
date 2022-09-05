package main

import (
	"fmt"
)

func main() {
	// if else Expression
	state := "idle"
	if state == "idle" {
		fmt.Println("standbay")
	} else if state == "running" {
		fmt.Println("running")
	} else {
		fmt.Println("stoped")
	}
	// sort statement
	if length := len("syarif"); length > 6 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama terlalu pendek")
	}
	// switch Expression
	// switch digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

	name := "Syarif"

	switch name {
	case "Syarif":
		fmt.Println("Hello Syarif")
	case "Al":
		fmt.Println("Hallo Al")
	default:
		fmt.Println("Hello, ")
	}

	// switch dengan sort statement
	switch len(name) > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}

	// switch tanpa kondisi
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu panjang")
	case panjang > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
