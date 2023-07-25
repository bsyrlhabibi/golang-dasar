package main

import "fmt"

func getFullname() (string, string) {
	return "Basyarul", "Habibi"

}
func main() {
	// Your main code heree
	firtsName, lastName := getFullname()
	fmt.Println(firtsName)
	fmt.Println(lastName)

	ignore, _ := getFullname()
	fmt.Println(ignore)
}
