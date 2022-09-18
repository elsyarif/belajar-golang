package main

import (
	"fmt"
	"sort"
)

// sort adalah package yang berisikan utilitas untuk proses pengurutan.
// agar bisa di urutkan, harus mengimplementasikan kontrak di interface sort.Interface

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Syarif", 29},
		{"Ahmad", 27},
		{"Dayat", 24},
		{"Mamat", 32},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
