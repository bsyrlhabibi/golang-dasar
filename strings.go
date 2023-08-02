package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Basyarul Habibi", "Habibi")) //true
	fmt.Println(strings.Contains("Basyarul Habibi", "Adel"))   //false

	fmt.Println(strings.Split("Basyarul Habibi", " "))

	fmt.Println(strings.ToLower("Basyarul Habibi"))
	fmt.Println(strings.ToUpper("Basyarul Habibi"))

	fmt.Println(strings.Trim("     Basyarul Habibi    ", " "))

	fmt.Println(strings.ReplaceAll("Bibi Bibi Bibi Joko", "Bibi", "Andi"))

}
