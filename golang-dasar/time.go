package main

import (
	"fmt"
	"time"
)

//package yang berisikan functionalitas untuk menagement waktu di golang

func main() {
	now := time.Now() // membuat waktu saat ini

	fmt.Println(now.Unix())
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2021, 9, 21, 9, 00, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	layout := "2006-01-02"                       // format tanggal
	parse, _ := time.Parse(layout, "2020-11-08") // parsing tanggal

	fmt.Println("parse:", parse)
}
