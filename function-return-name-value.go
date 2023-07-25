package main

import "fmt"

func getFullName2() (firstName string, middleName string, umur int) {
	firstName = "Basyarul"
	middleName = "Habibi"
	umur = 22
	return
}

func main() {
	// Your main code here
	firstName, middleName, umur := getFullName2() //nama tidak harus sama
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(umur)
}
