package main

import (
	"fmt"
	"reflect"
)

// bisa melihat struktur kode pada saat aplikasi berjalan, dengan menggunakan package reflect
// reflect sangat berguna ketika kita ingin membuat library yang general sehingga mudah untuk di gunakan

//type Sample struct {
//	Name string
//}

// Sample struct tag
type Sample struct {
	Name string `required:"true" max:"10"`
}

// mambuat validati
func Isvalid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Syarif"}

	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0) // untuk mengakses file ke 0

	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")

	fmt.Println(sampleType.NumField()) // untuk mengetahui jumlah field
	fmt.Println(structField.Name)

	// tag
	fmt.Println("==== Tag ====")
	fmt.Println(required)
	fmt.Println(max)

	// Validasi
	sample.Name = ""
	fmt.Println("==== Validasi ====")
	fmt.Println(Isvalid(sample))
}
