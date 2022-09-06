package main

import "fmt"

// mambuat function langsung di variable atau parameter tanpa harus membuat functin terlebih dahulu

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Anda di blackist")
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Syarif", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
