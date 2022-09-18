package main

import (
	"fmt"
	"regexp"
)

// utilitas untuk melakukan pencarian regular expression
// https://github.com/google/re2/wiki/Syntax
// https://golang.org/pkg/regexp/

func main() {
	var regex = regexp.MustCompile(`s([a-z])f`)

	fmt.Println(regex.MatchString("saf"))

	search := regex.FindAllString("saf sif suf sef sAf", 10)

	fmt.Println(search)
}
