package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`b([a-z])i`)
	// var regex *regexp.Regexp = regexp.MustCompile("b([a-z][a-z])i") //menambahkan jumlah kata

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("byi"))  // true
	fmt.Println(regex.MatchString("boyi")) // false
	fmt.Println(regex.MatchString("biYi"))

	search := regex.FindAllString("boi eka eko edo eto ek", -1)
	fmt.Println(search)
}
