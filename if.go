package main

import "fmt"

func main() {
	name := "Bibi"

	if name == "Bibi" {
		fmt.Println("Hello Bibi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	//short statement
	var lenght = len(name)
	if lenght > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
