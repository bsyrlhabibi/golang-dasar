package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Basyarul"
	names[1] = "Habibi"
	names[2] = "Muhammadun"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array langsung
	var values = [3]int{
		80,
		90,
		70,
	}
	fmt.Println(values)

	// function array
	fmt.Println(len(names)) //len untuk panjang array
	fmt.Println(len(values))
}
