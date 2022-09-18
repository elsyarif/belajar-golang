package main

import (
	"fmt"
	"math"
)

func main() {
	round := math.Round(33.75) // pembulatan kebawah atau keatas, sesaui dengan yang paling dekat
	flor := math.Floor(33.56)  // pembulatan kebawah
	ceil := math.Ceil(56.34)   // pembulatan keatas
	max := math.Max(12, 44)    // mengambil nilai paling besar
	min := math.Min(11, 9)     // mengambil nilai paling kecil

	fmt.Println("Round:", round)
	fmt.Println("Flor", flor)
	fmt.Println("Ceil", ceil)
	fmt.Println("Max", max)
	fmt.Println("Min", min)
}
