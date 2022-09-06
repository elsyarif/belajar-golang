package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	fmt.Println("==== Error Interface ===")
	hasil, err := pembagian(100, 10)

	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
