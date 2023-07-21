package main

import "fmt"

func main() {

	//Versi Expert
	// person := map[string]string

	var person map[string]string = map[string]string{
		"name":    "Bibi",
		"address": "Subang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["tittle"] = "Buku Go-Lang"
	book["author"] = "Bibi"
	book["wrong"] = "Ah salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
