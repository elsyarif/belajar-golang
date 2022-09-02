package main

import "fmt"

func main() {

	fmt.Println("Satu : ", 1)
	fmt.Println("dua : ", 2)
	fmt.Println("Tiga koma Lima : ", 3.5)

	var int8 int8
	int8 = 127
	// Tipe data Number int8 Nilai minumum -128 ~ nilai maximum 127
	fmt.Println(int8)

	var int16 int16
	int16 = 32767
	// Tipe data Number int16 Nilai minimum -32768 ~ nilai maximum 32767
	fmt.Println(int16)

	var int32 int32
	int32 = 2147483647
	// Tipe data number int32 nilai minimum -2147483648 ~ nilai maximum 2147483647
	fmt.Println(int32)

	var int64 int64
	int64 = 9223372036854775807
	// Tipe data number int64 nilai minimum -9223372036854775808 ~ nilai maximum 9223372036854775807
	fmt.Println(int64)

	var uint8 uint8
	uint8 = 255
	// Tipe data number unit8 nilai minimum 0 ~ 55
	fmt.Println(uint8)

	var uint16 uint16
	uint16 = 65535
	// Tipe data number uint16 nilai minimum 0 ~ 65535
	fmt.Println(uint16)

	var uint32 uint32
	uint32 = 4294967295
	// Tipe data number uint32 nilai minimum 0 ~ 4294967295
	fmt.Println(uint32)

	var uint64 uint64
	uint64 = 18446744073709551615
	// Tipe data number uint64 nilai minimum 0 ~ 18446744073709551615
	fmt.Println(uint64)

	var float32 float32
	float32 = 0.000001
	// tipe data number float nilai minimum 1.8x10^-38 ~ 3.4×10^38
	fmt.Println(float32)

	var float64 float64
	float64 = 0.000001
	// tipe data number float nilai minimum 1.8x10^-38 ~ 3.4×10^38
	fmt.Println(float64)
}
