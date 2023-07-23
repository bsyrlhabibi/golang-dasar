package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}

	// result := "Hello"
	// return result
	// return "Hello" //tidak harus langsung data nya
}
func main() {
	result := getHello("Bibi")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
