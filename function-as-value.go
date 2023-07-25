package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Your main code here
	sayGoodBye := getGoodBye
	result := sayGoodBye("Bibi")
	fmt.Println(result)
	fmt.Println(getGoodBye("Bibi"))
}
