package main

import "fmt"

func main() {
	name := "Bibi"

	switch name {
	case "Bibi":
		fmt.Println("Hello Bibi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Salah")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama terlalu panjang")
	case lenght > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
