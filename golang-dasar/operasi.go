package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	// operasi matematika  => 	augmented assigment
	// a = a + 10 => a += 10
	// a = a - 10 => a -= 10
	// a = a * 10 => a *= 10
	// a = a / 10 => a /= 10
	// a = a % 10 => a %= 10
	var i = 10
	i += 10

	fmt.Println(i)

	// Unery operator
	// ++ a = a + 1
	// -- a = a - 1
	// - Negative
	// + Positive
	// ! Boolean kebalikan
	var j = 1
	j++
	j++
	fmt.Println(j)
}
