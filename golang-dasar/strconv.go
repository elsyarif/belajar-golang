package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true") // mengubah string ker boolean

	if err == nil {
		fmt.Println("bool:", boolean)
	} else {
		fmt.Println("err:", err.Error())
	}
	float, _ := strconv.ParseFloat("1000234.23", 10) // mengubah string ke float

	fmt.Println("Float:", float)

	number, _ := strconv.ParseInt("1020120101", 10, 32) // mengubah string ke int64

	fmt.Println("number:", number)

	value := strconv.FormatInt(123423, 2) // base: 2 , merubah ke binary; base: 8, berbah ke octal; base: 16, merubah ke hexa decimal

	fmt.Println(value)

	valueInt, _ := strconv.Atoi("200000") // merubah string ke integer

	fmt.Println(valueInt)

	valueStr := strconv.Itoa(20000) // merubah integer ke string
	fmt.Println(valueStr)
}
