package main

import (
	"fmt"
	"golang-dasar/database"
)

func main() {
	result := database.GetDatabase()

	fmt.Println(result)
}
